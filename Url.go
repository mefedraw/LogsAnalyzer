package main

import (
	"bufio"
	"fmt"
	"net/http"
	"sync"
)

func AnalyzeUrl(url string) {

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка запроса:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ошибка: статус ответа", resp.StatusCode)
		return
	}

	logsAnalyzerUtil := NewLogAnalyzerUtil()
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
			Parse(string(line), logsAnalyzerUtil)
		}(line)
	}

	wg.Wait()

	if logsAnalyzerUtil.LogsNumber > 0 {
		fmt.Printf("Количество логов: %d, Средний размер ответа: %d\n",
			logsAnalyzerUtil.LogsNumber, logsAnalyzerUtil.ResponseSizeSum/logsAnalyzerUtil.LogsNumber)
	} else {
		fmt.Println("Логи не найдены или не обработаны")
	}
}
