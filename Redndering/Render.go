package Redndering

import "NginxLogsAnalyzer/LogsUtil"

type Render interface {
	Render(logsInfo *LogsUtil.LogAnalyzerUtil)
}
