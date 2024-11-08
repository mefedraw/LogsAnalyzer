package Redndering

import "NginxLogsAnalyzer/LogsUtil"

type MarkdownRenderer struct {
}

func newMarkdownRenderer() *MarkdownRenderer {
	return &MarkdownRenderer{}
}

func (mr *MarkdownRenderer) Render(logsInfo *LogsUtil.LogAnalyzerUtil) {

}
