package dataCollecting_test

import (
	"NginxLogsAnalyzer/dataCollecting"
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

var logData = `93.180.71.3 - - [17/May/2015:08:05:32 +0000] "GET /downloads/product_1 HTTP/1.1" 304 0 "-" "Debian APT-HTTP/1.3 (0.8.16~exp12ubuntu10.21)"
80.91.33.133 - - [17/May/2015:08:05:24 +0000] "GET /downloads/product_1 HTTP/1.1" 304 0 "-" "Debian APT-HTTP/1.3 (0.8.16~exp12ubuntu10.17)"
217.168.17.5 - - [17/May/2015:08:05:34 +0000] "GET /downloads/product_1 HTTP/1.1" 200 490 "-" "Debian APT-HTTP/1.3 (0.8.10.3)"
217.168.17.5 - - [17/May/2015:08:05:02 +0000] "GET /downloads/product_2 HTTP/1.1" 404 337 "-" "Debian APT-HTTP/1.3 (0.8.10.3)"
91.234.194.89 - - [17/May/2015:08:05:22 +0000] "GET /downloads/product_2 HTTP/1.1" 304 0 "-" "Debian APT-HTTP/1.3 (0.9.7.9)"
`

func TestLogDataCollector_CollectData_Success(t *testing.T) {
	fromDate := "2015-05-16"
	toDate := "2015-05-18"
	filter := ""

	collector := dataCollecting.NewLogDataCollector(fromDate, toDate, filter)
	reader := bufio.NewReader(bytes.NewBufferString(logData))

	err := collector.CollectData(reader)
	assert.NoError(t, err)

	assert.Equal(t, int64(5), collector.LogsInfo.LogsNumber)
	assert.Equal(t, int64(827), collector.LogsInfo.ResponseSizeSum)
	assert.Equal(t, int64(1), collector.LogsInfo.MostFrequentStatusCodes[200])
	assert.Equal(t, int64(1), collector.LogsInfo.MostFrequentStatusCodes[404])
	assert.Equal(t, int64(1), collector.LogsInfo.ErrorStatusCodeCount)
}

func TestLogDataCollector_CollectData_Filtered_Success(t *testing.T) {
	fromDate := "2015-05-16"
	toDate := "2015-05-18"
	filter := "/downloads/product_1"

	collector := dataCollecting.NewLogDataCollector(fromDate, toDate, filter)
	reader := bufio.NewReader(bytes.NewBufferString(logData))

	err := collector.CollectData(reader)
	assert.NoError(t, err)

	assert.Equal(t, int64(3), collector.LogsInfo.LogsNumber)
	assert.Equal(t, int64(490), collector.LogsInfo.ResponseSizeSum)
	assert.Equal(t, int64(1), collector.LogsInfo.MostFrequentStatusCodes[200])
	assert.Equal(t, int64(0), collector.LogsInfo.MostFrequentStatusCodes[404])
	assert.Equal(t, int64(0), collector.LogsInfo.ErrorStatusCodeCount)
}

func TestLogDataCollector_CollectData_OutOfDateRange(t *testing.T) {
	fromDate := "2015-05-18"
	toDate := "2015-05-19"
	filter := ""

	collector := dataCollecting.NewLogDataCollector(fromDate, toDate, filter)
	reader := bufio.NewReader(bytes.NewBufferString(logData))

	err := collector.CollectData(reader)
	assert.NoError(t, err)

	assert.Equal(t, int64(0), collector.LogsInfo.LogsNumber)
	assert.Equal(t, int64(0), collector.LogsInfo.ResponseSizeSum)
}

func TestLogDataCollector_CollectData_DateRangeAndFilter(t *testing.T) {
	fromDate := "2015-05-17"
	toDate := "2015-05-17"
	filter := "/downloads/product_2"

	collector := dataCollecting.NewLogDataCollector(fromDate, toDate, filter)
	reader := bufio.NewReader(bytes.NewBufferString(logData))

	err := collector.CollectData(reader)
	assert.NoError(t, err)

	assert.Equal(t, int64(2), collector.LogsInfo.LogsNumber)
	assert.Equal(t, int64(337), collector.LogsInfo.ResponseSizeSum)
	assert.Equal(t, int64(1), collector.LogsInfo.MostFrequentStatusCodes[404])
	assert.Equal(t, int64(1), collector.LogsInfo.ErrorStatusCodeCount)
}
