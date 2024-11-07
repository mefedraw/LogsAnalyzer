package Analyzer

import (
	"NginxLogsAnalyzer/LogsUtil"
	"NginxLogsAnalyzer/Parsing"
)

type Analyzer interface {
	Analyze(path string, parser Parsing.LogsParser) *LogsUtil.LogAnalyzerUtil
}
