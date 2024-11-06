package LogsUtil

import "sync"

type LogAnalyzerUtil struct {
	LogsNumber               int64
	MostRequestableResources map[string]int64
	MostFrequentStatusCodes  map[int64]int64
	ResponseSizeSum          int64
	Mu                       sync.Mutex
}

func NewLogAnalyzerUtil() *LogAnalyzerUtil {
	return &LogAnalyzerUtil{
		MostRequestableResources: make(map[string]int64),
		MostFrequentStatusCodes:  make(map[int64]int64),
	}
}
