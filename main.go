package main

import (
	"NginxLogsAnalyzer/Analyzing"
	"NginxLogsAnalyzer/BufferedSource"
	"NginxLogsAnalyzer/DataCollecting"
	"NginxLogsAnalyzer/FileModel"
	"NginxLogsAnalyzer/Rendering"
	"fmt"
)

func main() {
	filePath := "https://raw.githubusercontent.com/elastic/examples/master/Common%20Data%20Formats/nginx_logs/nginx_logs"
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	reader, _ := BufferedSource.NewHttpResponseReaderProvider().DataBufferWrap(filePath)
	collector := DataCollecting.NewLogDataCollector()
	_ = collector.CollectData(reader)
	analyzedData := Analyzing.NewNginxLogAnalyzer().Analyze(&collector.LogsInfo)
	fileData := FileModel.NewFileModelBuilder().
		SetFileName(filePath).
		SetFileAnalyzedData(*analyzedData).Build()
	renderedData := Rendering.NewMarkdownRenderer().Render(&fileData)
	//wg.Wait()
	fmt.Println(renderedData)
}
