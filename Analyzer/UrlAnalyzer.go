package Analyzer

import (
	"NginxLogsAnalyzer/LogsUtil"
	"NginxLogsAnalyzer/Parsing"
	"bufio"
	"fmt"
	"net/http"
	"sync"
)

type UrlAnalyzer struct{}

func NewUrlAnalyzer() *UrlAnalyzer {
	return &UrlAnalyzer{}
}

func (ua *UrlAnalyzer) Analyze(path string, parser Parsing.LogsParser) *LogsUtil.LogAnalyzerUtil {

	resp, err := http.Get(path)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return nil
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка: статус ответа", resp.StatusCode)
		return nil
	}

	logsAnalyzerUtil := LogsUtil.NewLogAnalyzerUtil()
	reader := bufio.NewReader(resp.Body)
	var wg sync.WaitGroup

	for {
		line, err := reader.ReadSlice('\n')
		if err != nil {
			break
		}
		wg.Add(1)
		go func(line []byte) {
			defer wg.Done()
			parser.ParseLine(string(line), logsAnalyzerUtil)
		}(line)
	}

	wg.Wait()

	if logsAnalyzerUtil.LogsNumber > 0 {
		fmt.Printf("Количество логов: %d, Средний размер ответа: %d\n",
			logsAnalyzerUtil.LogsNumber, logsAnalyzerUtil.ResponseSizeSum/logsAnalyzerUtil.LogsNumber)
	} else {
		fmt.Println("Логи не найдены или не обработаны")
	}

	return logsAnalyzerUtil
}
