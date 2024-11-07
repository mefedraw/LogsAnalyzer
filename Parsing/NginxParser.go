package Parsing

import (
	"NginxLogsAnalyzer/LogsUtil"
	"regexp"
	"strconv"
	"sync/atomic"
)

type NginxLogsParser struct {
}

func NewNginxLogsParser() *NginxLogsParser {
	return &NginxLogsParser{}
}

func (nlp *NginxLogsParser) ParseLine(line string, logAnalyzerUtil *LogsUtil.LogAnalyzerUtil) {
	re := regexp.MustCompile(`\[(\d{2}/\w{3}/\d{4}):.*?\] "(?:GET|HEAD|POST|PATCH) (/downloads/[\w\d_]+) HTTP/.*?" (\d{3}) (\d+)`)
	matches := re.FindStringSubmatch(line)

	// Проверка, что регулярное выражение нашло все необходимые группы
	if len(matches) < 5 {
		return
	}

	// Преобразование значений в int
	statusCode, _ := strconv.Atoi(matches[3])
	responseSize, _ := strconv.Atoi(matches[4])

	// Атомарные операции для общих переменных
	atomic.AddInt64(&logAnalyzerUtil.ResponseSizeSum, int64(responseSize))
	atomic.AddInt64(&logAnalyzerUtil.LogsNumber, 1)

	// Потокобезопасный доступ к картам с помощью мьютекса
	logAnalyzerUtil.Mu.Lock()
	logAnalyzerUtil.MostRequestableResources[matches[2]]++
	logAnalyzerUtil.MostFrequentStatusCodes[int64(statusCode)]++
	logAnalyzerUtil.Mu.Unlock()
}
