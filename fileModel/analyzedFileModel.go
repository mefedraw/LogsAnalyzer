package fileModel

import (
	LogsUtil "NginxLogsAnalyzer/logModel"
	"path/filepath"
	"strings"
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

	b.fileModel.FileName = getFileName(fileName)
	return b
}

func (b *FileModelBuilder) SetFromDate(fromDate string) *FileModelBuilder {
	b.fileModel.FromDate, _ = parseUserDate(fromDate)
	return b
}

func (b *FileModelBuilder) SetToDate(toDate string) *FileModelBuilder {
	b.fileModel.ToDate, _ = parseUserDate(toDate)
	return b
}

func (b *FileModelBuilder) SetFileAnalyzedData(data LogsUtil.LogAnalyzedData) *FileModelBuilder {
	b.fileModel.FileAnalyzedData = data
	return b
}

func (b *FileModelBuilder) Build() FileModel {
	return b.fileModel
}

func parseUserDate(userDate string) (time.Time, error) {
	if userDate == "" {
		return time.Time{}, nil
	}
	return time.Parse("2006-01-02", userDate)
}

func getFileName(path string) string {
	fileNameWithExt := filepath.Base(path)
	fileName := strings.TrimSuffix(fileNameWithExt, filepath.Ext(fileNameWithExt))
	return fileName
}
