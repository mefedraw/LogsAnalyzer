package main

import (
	"NginxLogsAnalyzer/Analyzer"
	"NginxLogsAnalyzer/Parsing"
)

func main() {
	filePath := "nginx_logs.txt"
	var analyzer = Analyzer.NewFileAnalyzer(filePath)
	var parser = Parsing.NewNginxLogsParser()
	analyzer.Analyze(parser)
}
