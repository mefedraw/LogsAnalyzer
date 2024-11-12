package main

import (
	"NginxLogsAnalyzer/analyzing"
	"NginxLogsAnalyzer/dataCollecting"
	"NginxLogsAnalyzer/fileModel"
	"NginxLogsAnalyzer/input"
	"NginxLogsAnalyzer/service"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (a *Application) Run() {
	inputData := input.NewUserInput()
	_ = inputData.Input()
	bufProvider, _ := service.NewFileReaderProviderService().GetReader(inputData.Path)
	reader, _ := bufProvider.DataBufferWrap(inputData.Path)
	collector := dataCollecting.NewLogDataCollector(inputData.FromDate, inputData.ToDate, inputData.Filter)
	_ = collector.CollectData(reader)
	analyzedData, _ := analyzing.NewNginxLogAnalyzer().Analyze(&collector.LogsInfo)

	fileData := fileModel.NewFileModelBuilder().
		SetFileName(inputData.Path).
		SetFromDate(inputData.FromDate).
		SetToDate(inputData.ToDate).
		SetFileAnalyzedData(*analyzedData).
		Build()
	render, _ := service.NewRenderService().GetRender(inputData.Format)
	render.Render(&fileData)
}
