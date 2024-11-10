package main

import (
	"NginxLogsAnalyzer/Analyzing"
	"NginxLogsAnalyzer/DataCollecting"
	"NginxLogsAnalyzer/FileModel"
	"NginxLogsAnalyzer/Input"
	"NginxLogsAnalyzer/Service"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Run() {
	inputData := Input.NewUserInput()
	inputData.Input()
	bufProvider, _ := Service.NewFileReaderProviderService().GetReader(inputData.Path)
	reader, _ := bufProvider.DataBufferWrap(inputData.Path)
	collector := DataCollecting.NewLogDataCollector(inputData.FromDate, inputData.ToDate, inputData.Filter)
	_ = collector.CollectData(reader)
	analyzedData := Analyzing.NewNginxLogAnalyzer().Analyze(&collector.LogsInfo)

	fileData := FileModel.NewFileModelBuilder().
		SetFileName(inputData.Path).
		SetFromDate(inputData.FromDate).
		SetToDate(inputData.ToDate).
		SetFileAnalyzedData(*analyzedData).
		Build()
	render, _ := Service.NewRenderService().GetRender(inputData.Format)
	render.Render(&fileData)
}
