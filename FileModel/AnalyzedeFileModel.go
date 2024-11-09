package FileModel

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

// Структура Builder для FileModel
type FileModelBuilder struct {
	fileModel FileModel
}

// Функция для создания нового билдера
func NewFileModelBuilder() *FileModelBuilder {
	return &FileModelBuilder{}
}

// Метод для задания имени файла
func (b *FileModelBuilder) SetFileName(fileName string) *FileModelBuilder {
	b.fileModel.FileName = fileName
	return b
}

// Метод для задания начальной даты
func (b *FileModelBuilder) SetFromDate(fromDate time.Time) *FileModelBuilder {
	b.fileModel.FromDate = fromDate
	return b
}

// Метод для задания конечной даты
func (b *FileModelBuilder) SetToDate(toDate time.Time) *FileModelBuilder {
	b.fileModel.ToDate = toDate
	return b
}

// Метод для задания данных анализа
func (b *FileModelBuilder) SetFileAnalyzedData(data LogsUtil.LogAnalyzedData) *FileModelBuilder {
	b.fileModel.FileAnalyzedData = data
	return b
}

// Метод для построения и возврата FileModel
func (b *FileModelBuilder) Build() FileModel {
	return b.fileModel
}
