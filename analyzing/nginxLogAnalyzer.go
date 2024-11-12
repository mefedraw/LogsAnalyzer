package analyzing

import (
	"NginxLogsAnalyzer/logModel"
	"sort"
)

type NginxLogAnalyzer struct{}

func NewNginxLogAnalyzer() *NginxLogAnalyzer {
	return &NginxLogAnalyzer{}
}

func (nla *NginxLogAnalyzer) Analyze(logsCollectedData *logModel.LogDataCollectUtil) (*logModel.LogAnalyzedData, error) {
	analyzedDataBuilder := logModel.NewLogAnalyzedDataBuilder()
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
		SetUniqueIpCount(nla.GetUniqueIpCount(logsCollectedData.Ips)).
		SetErrorStatusCodePercentage(
			nla.CalcErrorStatusCodePercentage(logsCollectedData.LogsNumber, logsCollectedData.ErrorStatusCodeCount),
		).Build()

	return &analyzedData, nil
}

func (nla *NginxLogAnalyzer) GetTop3StatusCodes(statusCodes map[int64]int64) []logModel.CodeCountTuple {
	var tuples []logModel.CodeCountTuple
	for code, count := range statusCodes {
		tuples = append(tuples, logModel.CodeCountTuple{Code: code, Count: count})
	}

	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].Count > tuples[j].Count
	})

	var top3 []logModel.CodeCountTuple
	for i := 0; i < 3 && i < len(tuples); i++ {
		top3 = append(top3, tuples[i])
	}

	return top3
}

func (nla *NginxLogAnalyzer) GetTop3ServerResources(statusCodes map[string]int64) []logModel.ResourceCount {
	var tuples []logModel.ResourceCount
	for response, count := range statusCodes {
		tuples = append(tuples, logModel.ResourceCount{Resource: response, Count: count})
	}

	sort.Slice(tuples, func(i, j int) bool {
		return tuples[i].Count > tuples[j].Count
	})

	var top3 []logModel.ResourceCount
	for i := 0; i < 3 && i < len(tuples); i++ {
		top3 = append(top3, tuples[i])
	}

	return top3
}

func (nla *NginxLogAnalyzer) CalcAverageServerResponseSize(logsNum, serverResponseSizeSum int64) int64 {
	if logsNum == 0 {
		return 0
	}
	return (serverResponseSizeSum) / (logsNum)
}

func (nla *NginxLogAnalyzer) GetUniqueIpCount(ips map[string]int64) int64 {
	return int64(len(ips))
}

func (nla *NginxLogAnalyzer) CalcErrorStatusCodePercentage(logsNum, errorStatusCodeCount int64) float64 {
	return (float64(errorStatusCodeCount) / float64(logsNum)) * 100
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
