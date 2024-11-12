package parsing

import "NginxLogsAnalyzer/logModel"

type LogsParser interface {
	ParseLine(path string, logAnalyzerUtil *logModel.LogDataCollectUtil)
}
