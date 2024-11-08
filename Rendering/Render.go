package Rendering

import LogsUtil "NginxLogsAnalyzer/LogModel"

type Render interface {
	Render(analyzedData *LogsUtil.LogAnalyzedData)
}
