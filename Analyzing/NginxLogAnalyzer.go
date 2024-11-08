package Analyzing

import (
	"NginxLogsAnalyzer/LogModel"
	"sort"
)

type NginxLogAnalyzer struct{}

func NewNginxLogAnalyzer() *NginxLogAnalyzer {
	return &NginxLogAnalyzer{}
}

func (nla *NginxLogAnalyzer) Analyze(logsCollectedData *LogsUtil.LogDataCollectUtil) *LogsUtil.LogAnalyzedData {
	analyzedDataBuilder := LogsUtil.NewLogAnalyzedDataBuilder()
	analyzedDataBuilder = analyzedDataBuilder.SetTotalRequests(logsCollectedData.LogsNumber)

	top3StatusCodes := nla.GetTop3StatusCodes(logsCollectedData.MostFrequentStatusCodes)
	for _, codeCountTuple := range top3StatusCodes {
		analyzedDataBuilder = analyzedDataBuilder.AddFrequentStatusCode(codeCountTuple.Code, codeCountTuple.Count)
	}

	top3Resources := nla.GetTop3ServerResources(logsCollectedData.MostRequestableResources)
	for _, resourceCountTuple := range top3Resources {
		analyzedDataBuilder = analyzedDataBuilder.AddFrequentResource(resourceCountTuple.Resource, resourceCountTuple.Count)
	}

	analyzedData := analyzedDataBuilder.
		SetAverageResponseSize(
			nla.CalcAverageServerResponseSize(logsCollectedData.LogsNumber, logsCollectedData.ResponseSizeSum),
		).
		SetResponseSize95Percentile(
			nla.Calc95PercentileServerResponseSize(logsCollectedData.AllServerResponses),
		).
		Build()

	return &analyzedData
}

func (nla *NginxLogAnalyzer) GetTop3StatusCodes(statusCodes map[int64]int64) [3]LogsUtil.CodeCountTuple {
	var tuples []LogsUtil.CodeCountTuple
	for code, count := range statusCodes {
		tuples = append(tuples, LogsUtil.CodeCountTuple{Code: code, Count: count})
	}

	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].Count > tuples[j].Count
	})

	var top3 [3]LogsUtil.CodeCountTuple
	for i := 0; i < 3 && i < len(tuples); i++ {
		top3[i] = tuples[i]
	}

	return top3
}

func (nla *NginxLogAnalyzer) GetTop3ServerResources(statusCodes map[string]int64) [3]LogsUtil.ResourceCount {
	var tuples []LogsUtil.ResourceCount
	for response, count := range statusCodes {
		tuples = append(tuples, LogsUtil.ResourceCount{Resource: response, Count: count})
	}

	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].Count > tuples[j].Count
	})

	var top3 [3]LogsUtil.ResourceCount
	for i := 0; i < 3 && i < len(tuples); i++ {
		top3[i] = tuples[i]
	}

	return top3
}

func (nla *NginxLogAnalyzer) CalcAverageServerResponseSize(logsNum, serverResponseSizeSum int64) int64 {
	if logsNum == 0 {
		return 500
	}
	return (serverResponseSizeSum) / (logsNum)
}

func (nla *NginxLogAnalyzer) Calc95PercentileServerResponseSize(responseSizes []int64) int64 {
	if len(responseSizes) == 0 {
		return 0
	}

	sort.Slice(responseSizes, func(i, j int) bool {
		return responseSizes[i] < responseSizes[j]
	})

	index := int(float64(len(responseSizes)) * 0.95)
	if index >= len(responseSizes) {
		index = len(responseSizes) - 1
	}

	return responseSizes[index]
}
