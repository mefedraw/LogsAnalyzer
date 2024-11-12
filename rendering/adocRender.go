package rendering

import (
	"NginxLogsAnalyzer/fileModel"
	"fmt"
	"net/http"
)

type AdocRender struct{}

func NewAdocRender() *AdocRender {
	return &AdocRender{}
}

func (ar *AdocRender) Render(file *fileModel.FileModel) {
	fmt.Print(ar.BuildReportString(file))
}

func (ar *AdocRender) BuildReportString(file *fileModel.FileModel) string {
	data := file.FileAnalyzedData

	fromDate := "-"
	if !file.FromDate.IsZero() {
		fromDate = file.FromDate.Format("02.01.2006")
	}

	toDate := "-"
	if !file.ToDate.IsZero() {
		toDate = file.ToDate.Format("02.01.2006")
	}

	result := "== Общая информация\n\n"
	result += "|===\n"
	result += "| Метрика                | Значение\n"
	result += "| Файл(-ы)               | " + file.FileName + "\n"
	result += "| Начальная дата         | " + fromDate + "\n"
	result += "| Конечная дата          | " + toDate + "\n"
	result += fmt.Sprintf("| Количество запросов    | %d\n", data.TotalRequests)
	result += fmt.Sprintf("| Средний размер ответа  | %db\n", data.AverageResponseSize)
	result += fmt.Sprintf("| 95p размера ответа     | %db\n", data.ResponseSize95Percentile)
	result += fmt.Sprintf("| Процент ошибок         | %.2f%%\n", data.ErrorStatusCodePercentage)
	result += fmt.Sprintf("| Кол-во уникальных Ip   | %d\n", data.UniqueIpCount)
	result += "|===\n\n"

	result += "== Запрашиваемые ресурсы\n\n"
	result += "|===\n"
	result += "| Ресурс                 | Количество\n"
	for _, resource := range data.MostFrequentResources {
		result += fmt.Sprintf("| %-22s | %d\n", resource.Resource, resource.Count)
	}
	result += "|===\n\n"

	result += "== Коды ответа\n\n"
	result += "|===\n"
	result += "| Код  | Имя                  | Количество\n"
	for _, code := range data.MostFrequentStatusCodes {
		result += fmt.Sprintf("| %-4d | %-20s | %d\n", code.Code, http.StatusText(int(code.Code)), code.Count)
	}
	result += "|===\n"

	return result
}
