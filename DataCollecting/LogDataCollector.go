package DataCollecting

import (
	"NginxLogsAnalyzer/DataReading"
	"NginxLogsAnalyzer/LogModel"
	"NginxLogsAnalyzer/Parsing"
	"bufio"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
)

type LogDataCollector struct {
	LogsInfo LogsUtil.LogDataCollectUtil
}

func NewLogDataCollector() *LogDataCollector {
	return &LogDataCollector{
		LogsInfo: *LogsUtil.NewLogDataCollectUtil(),
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
				var matches = parser.ParseLine(line)
				ldc.updateInfo(*matches)
			}
		}(i)
	}
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
