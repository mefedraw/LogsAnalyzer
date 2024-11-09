package LogsUtil

import "sync"

type LogDataCollectUtil struct {
	LogsNumber               int64
	MostRequestableResources map[string]int64
	MostFrequentStatusCodes  map[int64]int64
	AllServerResponses       []int64
	ResponseSizeSum          int64
	ErrorStatusCodeCount     int64
	Mu                       sync.Mutex
}

func NewLogDataCollectUtil() *LogDataCollectUtil {
	return &LogDataCollectUtil{
		MostRequestableResources: make(map[string]int64),
		MostFrequentStatusCodes:  make(map[int64]int64),
		AllServerResponses:       make([]int64, 0),
	}
}
