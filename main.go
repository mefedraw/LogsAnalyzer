package main

import (
	"NginxLogsAnalyzer/AnalyzedFile"
	"NginxLogsAnalyzer/Analyzing"
	"NginxLogsAnalyzer/BufferedSource"
	"NginxLogsAnalyzer/DataCollecting"
	"NginxLogsAnalyzer/Rendering"
	"fmt"
)

func main() {
	filePath := "nginx_logs.txt"
	bufioR, _ := BufferedSource.NewFileReaderProvider().DataBufferWrap(filePath)
	dataCollector := DataCollecting.NewLogDataCollector()

	_ = dataCollector.CollectData(bufioR)
	analyzedData := Analyzing.NewNginxLogAnalyzer().Analyze(&dataCollector.LogsInfo)
	analyzeFile := AnalyzedFile.NewFileModelBuilder().SetFileName(filePath).SetFileAnalyzedData(*analyzedData).Build()
	renderedInfo := Rendering.NewMarkdownRenderer().Render(&analyzeFile)
	fmt.Println(renderedInfo)
}
