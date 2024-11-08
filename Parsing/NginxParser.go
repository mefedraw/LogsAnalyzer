package Parsing

import (
	"NginxLogsAnalyzer/LogModel"
	"regexp"
)

type NginxLogsParser struct {
}

func NewNginxLogsParser() *NginxLogsParser {
	return &NginxLogsParser{}
}

func (nlp *NginxLogsParser) ParseLine(line string, logAnalyzerUtil *LogsUtil.LogDataCollectUtil) *[]string {
	re := regexp.MustCompile(`\[(\d{2}/\w{3}/\d{4}):.*?\] "(?:GET|HEAD|POST|PATCH) (/downloads/[\w\d_]+) HTTP/.*?" (\d{3}) (\d+)`)
	matches := re.FindStringSubmatch(line)

	if len(matches) < 5 {
		return nil
	}
	return &matches
}
