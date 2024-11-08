package Rendering

import (
	"NginxLogsAnalyzer/AnalyzedFile"
	"fmt"
	"net/http"
)

type MarkdownRenderer struct{}

func NewMarkdownRenderer() *MarkdownRenderer {
	return &MarkdownRenderer{}
}

func (mdr *MarkdownRenderer) Render(file *AnalyzedFile.FileModel) string {
	data := file.FileAnalyzedData

	// Общая информация
	result := "#### Общая информация\n\n"
	result += "| Метрика                | Значение     |\n"
	result += "|------------------------|--------------|\n"
	result += fmt.Sprintf("| Файл(-ы)               | %s           |\n", file.FileName)
	result += fmt.Sprintf("| Начальная дата         | %s           |\n", file.FromDate.Format("02.01.2006"))
	result += fmt.Sprintf("| Конечная дата          | %s           |\n", file.ToDate.Format("02.01.2006"))
	result += fmt.Sprintf("| Количество запросов    | %d           |\n", data.TotalRequests)
	result += fmt.Sprintf("| Средний размер ответа  | %db          |\n", data.AverageResponseSize)
	result += fmt.Sprintf("| 95p размера ответа     | %db          |\n", data.ResponseSize95Percentile)

	// Запрашиваемые ресурсы
	result += "\n#### Запрашиваемые ресурсы\n\n"
	result += "| Ресурс                 | Количество   |\n"
	result += "|------------------------|--------------|\n"
	for _, resource := range data.MostFrequentResources {
		result += fmt.Sprintf("| %-22s | %d           |\n", resource.Resource, resource.Count)
	}

	// Коды ответа
	result += "\n#### Коды ответа\n\n"
	result += "| Код  | Имя                  | Количество   |\n"
	result += "|------|----------------------|--------------|\n"
	for _, code := range data.MostFrequentStatusCodes {
		result += fmt.Sprintf("| %-4d | %-20s | %d           |\n", code.Code, http.StatusText(int(code.Code)), code.Count)
	}

	return result
}
