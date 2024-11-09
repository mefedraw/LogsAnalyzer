﻿package AnalyzedFile

import (
	LogsUtil "NginxLogsAnalyzer/LogModel"
	"time"
)

type FileModel struct {
	FileName         string
	FromDate         time.Time
	ToDate           time.Time
	FileAnalyzedData LogsUtil.LogAnalyzedData
}

type FileModelBuilder struct {
	fileModel FileModel
}

func NewFileModelBuilder() *FileModelBuilder {
	return &FileModelBuilder{}
}

func (b *FileModelBuilder) SetFileName(fileName string) *FileModelBuilder {
	b.fileModel.FileName = fileName
	return b
}

func (b *FileModelBuilder) SetFromDate(fromDate time.Time) *FileModelBuilder {
	b.fileModel.FromDate = fromDate
	return b
}

func (b *FileModelBuilder) SetToDate(toDate time.Time) *FileModelBuilder {
	b.fileModel.ToDate = toDate
	return b
}

func (b *FileModelBuilder) SetFileAnalyzedData(data LogsUtil.LogAnalyzedData) *FileModelBuilder {
	b.fileModel.FileAnalyzedData = data
	return b
}

func (b *FileModelBuilder) Build() FileModel {
	return b.fileModel
}
