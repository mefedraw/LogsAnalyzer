package Rendering

import (
	"NginxLogsAnalyzer/FileModel"
	"fmt"
	"net/http"
)

type MarkdownRender struct {
}

func NewMarkdownRenderer() *MarkdownRender {
	return &MarkdownRender{}
}

func (mdr *MarkdownRender) Render(file *FileModel.FileModel) {
	renderInfo := mdr.BuildReportString(file)
	fmt.Print(renderInfo)
}

func (mdr *MarkdownRender) BuildReportString(file *FileModel.FileModel) string {
	data := file.FileAnalyzedData

	fromDate := "-"
	if !file.FromDate.IsZero() {
		fromDate = file.FromDate.Format("02.01.2006")
	}

	toDate := "-"
	if !file.ToDate.IsZero() {
		toDate = file.ToDate.Format("02.01.2006")
	}

	result := "#### Общая информация\n\n"
	result += "| Метрика                | Значение     |\n"
	result += "|------------------------|--------------|\n"
	result += fmt.Sprintf("| Файл(-ы)               | %s           |\n", file.FileName)
	result += fmt.Sprintf("| Начальная дата         | %s           |\n", fromDate)
	result += fmt.Sprintf("| Конечная дата          | %s           |\n", toDate)
	result += fmt.Sprintf("| Количество запросов    | %d           |\n", data.TotalRequests)
	result += fmt.Sprintf("| Средний размер ответа  | %db          |\n", data.AverageResponseSize)
	result += fmt.Sprintf("| 95p размера ответа     | %db          |\n", data.ResponseSize95Percentile)
	result += fmt.Sprintf("| Процент ошибок         | %.2f%%       |\n", data.ErrorStatusCodePercentage)
	result += fmt.Sprintf("| Кол-во уникальных Ip   | %d       	  |\n", data.UniqueIpCount)

	result += "\n#### Запрашиваемые ресурсы\n\n"
	result += "| Ресурс                 | Количество   |\n"
	result += "|------------------------|--------------|\n"
	for _, resource := range data.MostFrequentResources {
		result += fmt.Sprintf("| %-22s | %d           |\n", resource.Resource, resource.Count)
	}

	result += "\n#### Коды ответа\n\n"
	result += "| Код  | Имя                  | Количество   |\n"
	result += "|------|----------------------|--------------|\n"
	for _, code := range data.MostFrequentStatusCodes {
		result += fmt.Sprintf("| %-4d | %-20s | %d           |\n", code.Code, http.StatusText(int(code.Code)), code.Count)
	}

	return result
}
