package main

import (
	"NginxLogsAnalyzer/Analyzer"
)

func main() {
	filePath := "nginx_logs.txt"
	var analyzer = Analyzer.NewFileAnalyzer()
	analyzer.Analyze(filePath)
}
