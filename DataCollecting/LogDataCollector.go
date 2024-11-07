package DataCollecting

import (
	"NginxLogsAnalyzer/DataReading"
	"NginxLogsAnalyzer/LogsUtil"
	"NginxLogsAnalyzer/Parsing"
	"bufio"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
)

type LogDataCollector struct {
	LogsInfo LogsUtil.LogAnalyzerUtil
}

func NewLogDataCollector(filePath string) *LogDataCollector {
	return &LogDataCollector{
		LogsInfo: *LogsUtil.NewLogAnalyzerUtil(),
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
				var matches = parser.ParseLine(line, &ldc.LogsInfo)
				ldc.UpdateInfo(*matches)
			}
		}(i)
	}
}

func (ldc *LogDataCollector) UpdateInfo(matches []string) {
	statusCode, _ := strconv.Atoi(matches[3])
	responseSize, _ := strconv.Atoi(matches[4])

	atomic.AddInt64(&ldc.LogsInfo.ResponseSizeSum, int64(responseSize))
	atomic.AddInt64(&ldc.LogsInfo.LogsNumber, 1)

	ldc.LogsInfo.Mu.Lock()
	ldc.LogsInfo.MostRequestableResources[matches[2]]++
	ldc.LogsInfo.MostFrequentStatusCodes[int64(statusCode)]++
	ldc.LogsInfo.Mu.Unlock()
}