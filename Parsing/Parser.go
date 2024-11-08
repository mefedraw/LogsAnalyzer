package Parsing

import "NginxLogsAnalyzer/LogModel"

type LogsParser interface {
	ParseLine(path string, logAnalyzerUtil *LogsUtil.LogDataCollectUtil)
}
