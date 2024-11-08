package Analyzing

import "NginxLogsAnalyzer/LogModel"

type Analyzer interface {
	Analyze(logsCollectedData *LogsUtil.LogDataCollectUtil) *LogsUtil.LogAnalyzedData
}
