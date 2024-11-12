package analyzing_test

import (
	"NginxLogsAnalyzer/analyzing"
	"testing"

	"NginxLogsAnalyzer/logModel"
	"github.com/stretchr/testify/assert"
)

func TestNginxLogAnalyzer_Analyze_Success(t *testing.T) {
	logsCollectedData := logModel.LogDataCollectUtil{
		LogsNumber:              100,
		ResponseSizeSum:         10000,
		MostFrequentStatusCodes: map[int64]int64{200: 80, 404: 15, 500: 5},
		MostRequestableResources: map[string]int64{
			"/home": 50, "/about": 30, "/contact": 20,
		},
		ErrorStatusCodeCount: 20,
		Ips:                  map[string]int64{"192.168.1.1": 1, "192.168.1.2": 1},
		AllServerResponses:   []int64{500, 1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000},
	}

	analyzer := analyzing.NewNginxLogAnalyzer()
	analyzedData, err := analyzer.Analyze(&logsCollectedData)

	assert.NoError(t, err)
	assert.Equal(t, int64(100), analyzedData.TotalRequests)
	assert.Equal(t, int64(2), analyzedData.UniqueIpCount)
	assert.Equal(t, int64(100), analyzedData.AverageResponseSize)
	assert.Equal(t, int64(9000), analyzedData.ResponseSize95Percentile)
	assert.Equal(t, float64(20), analyzedData.ErrorStatusCodePercentage)

	expectedStatusCodes := []logModel.CodeCountTuple{
		{Code: 200, Count: 80},
		{Code: 404, Count: 15},
		{Code: 500, Count: 5},
	}
	expectedResources := []logModel.ResourceCount{
		{Resource: "/home", Count: 50},
		{Resource: "/about", Count: 30},
		{Resource: "/contact", Count: 20},
	}
	assert.Equal(t, expectedStatusCodes, analyzedData.MostFrequentStatusCodes)
	assert.Equal(t, expectedResources, analyzedData.MostFrequentResources)
}

func TestNginxLogAnalyzer_CalcErrorStatusCodePercentage_Success(t *testing.T) {
	analyzer := analyzing.NewNginxLogAnalyzer()
	percentage := analyzer.CalcErrorStatusCodePercentage(100, 20)
	assert.Equal(t, 20.0, percentage)
}
