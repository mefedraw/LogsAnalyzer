package main

import (
	"NginxLogsAnalyzer/Errors/UserInputError"
	"NginxLogsAnalyzer/analyzing"
	"NginxLogsAnalyzer/dataCollecting"
	"NginxLogsAnalyzer/fileModel"
	"NginxLogsAnalyzer/input"
	"NginxLogsAnalyzer/service"
	"fmt"
	"log/slog"
	"os"
)

type Application struct {
	logger *slog.Logger
}

func NewApplication() *Application {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
	return &Application{
		logger: logger,
	}
}

func (a *Application) handleError(err error) {
	if err != nil {
		if userErr, ok := err.(*UserInputError.ErrUserInput); ok {
			fmt.Println(userErr.Error())
		} else {
			a.logger.Error("System Error:", "error", err.Error())
		}
		os.Exit(1)
	}
}

func (a *Application) Run() {
	inputData := input.NewUserInput()
	err := inputData.Input()
	a.handleError(err)
	bufProvider, err := service.NewFileReaderProviderService().GetReader(inputData.Path)
	a.handleError(err)
	reader, err := bufProvider.DataBufferWrap(inputData.Path)
	a.handleError(err)
	collector := dataCollecting.NewLogDataCollector(inputData.FromDate, inputData.ToDate, inputData.Filter)
	err = collector.CollectData(reader)
	a.handleError(err)
	analyzedData, err := analyzing.NewNginxLogAnalyzer().Analyze(&collector.LogsInfo)
	a.handleError(err)

	fileData := fileModel.NewFileModelBuilder().
		SetFileName(inputData.Path).
		SetFromDate(inputData.FromDate).
		SetToDate(inputData.ToDate).
		SetFileAnalyzedData(*analyzedData).
		Build()
	render, err := service.NewRenderService().GetRender(inputData.Format)
	a.handleError(err)
	render.Render(&fileData)
}
