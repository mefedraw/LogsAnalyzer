﻿package main

import (
	"NginxLogsAnalyzer/Analyzing"
	"NginxLogsAnalyzer/BufferedSource"
	"NginxLogsAnalyzer/DataCollecting"
	"NginxLogsAnalyzer/FileModel"
	"NginxLogsAnalyzer/Redndering"
	"fmt"
)

func main() {
	filePath := "nginx_logs.txt"
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	reader, _ := BufferedSource.NewFileReaderProvider().DataBufferWrap(filePath)
	collector := DataCollecting.NewLogDataCollector()
	_ = collector.CollectData(reader)
	analyzedData := Analyzing.NewNginxLogAnalyzer().Analyze(&collector.LogsInfo)
	fileData := FileModel.NewFileModelBuilder().
		SetFileName(filePath).
		SetFileAnalyzedData(*analyzedData).Build()
	renderedData := Redndering.NewMarkdownRenderer().Render(&fileData)
	//wg.Wait()
	fmt.Println(renderedData)
}
