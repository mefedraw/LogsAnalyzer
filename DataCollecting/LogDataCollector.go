package DataCollecting

import (
	"NginxLogsAnalyzer/DataReading"
	"NginxLogsAnalyzer/LogModel"
	"NginxLogsAnalyzer/Parsing"
	"bufio"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

type LogDataCollector struct {
	LogsInfo LogsUtil.LogDataCollectUtil
	FromDate time.Time
	ToDate   time.Time
	Filter   string
}

func NewLogDataCollector(fromDate, toDate, filter string) *LogDataCollector {
	FromDate, _ := parseUserDate(fromDate)
	ToDate, _ := parseUserDate(toDate)
	return &LogDataCollector{
		LogsInfo: *LogsUtil.NewLogDataCollectUtil(),
		FromDate: FromDate,
		ToDate:   ToDate,
		Filter:   filter,
	}
}

func (ldc *LogDataCollector) CollectData(reader *bufio.Reader) error {
	lines := make(chan string)
	var dataReader = DataReading.NewBufioDataReader()
	go func() {
		err := dataReader.ReadBuffer(reader, lines)
		if err != nil {
			return
		}
	}()

	var wg sync.WaitGroup
	ldc.collectLines(lines, &wg)
	wg.Wait()

	return nil
}

func (ldc *LogDataCollector) collectLines(lines chan string, wg *sync.WaitGroup) {
	var parser = Parsing.NewNginxLogsParser()
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for line := range lines {
				if ldc.Filter == "" || strings.Contains(line, ldc.Filter) {
					var matches = *parser.ParseLine(line)
					if matches == nil {
						continue
					}
					date, _ := parseLogDate(matches[2])
					if !ldc.ToDate.IsZero() || !ldc.FromDate.IsZero() {
						if (ldc.FromDate.IsZero() || ldc.FromDate.Before(date) || ldc.FromDate.Equal(date)) &&
							(ldc.ToDate.IsZero() || ldc.ToDate.After(date) || ldc.ToDate.Equal(date)) {
							ldc.updateInfo(matches)
						}
					} else {
						ldc.updateInfo(matches)
					}
				}
			}
		}(i)
	}
}

func parseLogDate(logDate string) (time.Time, error) {
	return time.Parse("02/Jan/2006", logDate)
}

func (ldc *LogDataCollector) updateInfo(matches []string) {
	statusCode, _ := strconv.Atoi(matches[4])
	responseSize, _ := strconv.Atoi(matches[5])

	atomic.AddInt64(&ldc.LogsInfo.ResponseSizeSum, int64(responseSize))
	atomic.AddInt64(&ldc.LogsInfo.LogsNumber, 1)

	ldc.LogsInfo.Mu.Lock()
	ldc.LogsInfo.MostRequestableResources[matches[3]]++
	ldc.LogsInfo.MostFrequentStatusCodes[int64(statusCode)]++
	if statusCode >= 400 && statusCode <= 599 {
		ldc.LogsInfo.ErrorStatusCodeCount++
	}
	ldc.LogsInfo.Ips[matches[1]]++
	ldc.LogsInfo.AllServerResponses = append(ldc.LogsInfo.AllServerResponses, int64(responseSize))
	ldc.LogsInfo.Mu.Unlock()

}

func parseUserDate(userDate string) (time.Time, error) {
	if userDate == "" {
		return time.Time{}, nil
	}
	return time.Parse("2006-01-02", userDate)
}
