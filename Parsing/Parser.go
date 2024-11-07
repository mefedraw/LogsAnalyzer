package Parsing

import "NginxLogsAnalyzer/LogsUtil"

type LogsParser interface {
	ParseLine(path string, logAnalyzerUtil *LogsUtil.LogAnalyzerUtil)
}
