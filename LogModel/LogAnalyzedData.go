package LogsUtil

type LogAnalyzedData struct {
	TotalRequests             int64
	MostFrequentResources     []ResourceCount
	MostFrequentStatusCodes   []CodeCountTuple
	AverageResponseSize       int64
	ResponseSize95Percentile  int64
	ErrorStatusCodePercentage float64
	UniqueIpCount             int64
}

func NewLogAnalyzedData() *LogAnalyzedData {
	return &LogAnalyzedData{}
}

type LogAnalyzedDataBuilder struct {
	logAnalyzedData LogAnalyzedData
}

func NewLogAnalyzedDataBuilder() *LogAnalyzedDataBuilder {
	return &LogAnalyzedDataBuilder{
		logAnalyzedData: *NewLogAnalyzedData(),
	}
}

func (builder *LogAnalyzedDataBuilder) SetTotalRequests(totalRequests int64) *LogAnalyzedDataBuilder {
	builder.logAnalyzedData.TotalRequests = totalRequests
	return builder
}

func (builder *LogAnalyzedDataBuilder) AddFrequentResource(frequentResource string, resourceCount int64) *LogAnalyzedDataBuilder {
	builder.logAnalyzedData.MostFrequentResources = append(builder.logAnalyzedData.MostFrequentResources, ResourceCount{frequentResource, resourceCount})
	return builder
}

func (builder *LogAnalyzedDataBuilder) AddFrequentStatusCode(frequentStatusCode int64, statusCodeCount int64) *LogAnalyzedDataBuilder {
	builder.logAnalyzedData.MostFrequentStatusCodes = append(builder.logAnalyzedData.MostFrequentStatusCodes, CodeCountTuple{frequentStatusCode, statusCodeCount})
	return builder
}

func (builder *LogAnalyzedDataBuilder) SetAverageResponseSize(averageResponseSize int64) *LogAnalyzedDataBuilder {
	builder.logAnalyzedData.AverageResponseSize = averageResponseSize
	return builder
}

func (builder *LogAnalyzedDataBuilder) SetResponseSize95Percentile(responseSize95Percentile int64) *LogAnalyzedDataBuilder {
	builder.logAnalyzedData.ResponseSize95Percentile = responseSize95Percentile
	return builder
}

func (builder *LogAnalyzedDataBuilder) SetErrorStatusCodePercentage(errorStatusCodePercentage float64) *LogAnalyzedDataBuilder {
	builder.logAnalyzedData.ErrorStatusCodePercentage = errorStatusCodePercentage
	return builder
}

func (builder *LogAnalyzedDataBuilder) SetUniqueIpCount(uniqueIpCount int64) *LogAnalyzedDataBuilder {
	builder.logAnalyzedData.UniqueIpCount = uniqueIpCount
	return builder
}

func (builder *LogAnalyzedDataBuilder) Build() LogAnalyzedData {
	return builder.logAnalyzedData
}
