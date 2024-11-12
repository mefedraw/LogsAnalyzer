package analyzing

import "NginxLogsAnalyzer/logModel"

type Analyzer interface {
	Analyze(logsCollectedData *logModel.LogDataCollectUtil) *logModel.LogAnalyzedData
}
