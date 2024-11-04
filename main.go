package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"sync"
	"sync/atomic"
)

type LogAnalyzerUtil struct {
	LogsNumber               int64
	MostRequestableResources map[string]int64
	mostFrequentStatusCodes  map[int64]int64
	ResponseSizeSum          int64
	mu                       sync.Mutex
}

func NewLogAnalyzerUtil() *LogAnalyzerUtil {
	return &LogAnalyzerUtil{
		MostRequestableResources: make(map[string]int64),
		mostFrequentStatusCodes:  make(map[int64]int64),
	}
}

func main() {
	filePath := "nginx_logs.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Ошибка получения информации о файле:", err)
		return
	}

	fileSize := fileInfo.Size()
	numWorkers := 4 // Количество горутин
	chunkSize := fileSize / int64(numWorkers)

	logsAnalyzerUtil := NewLogAnalyzerUtil()
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)

		// Вычисляем начало и конец каждого чанка
		start := int64(i) * chunkSize
		end := start + chunkSize
		if i == numWorkers-1 {
			end = fileSize // последний чанк до конца файла
		}

		go func(start, end int64) {
			defer wg.Done()
			processChunk(filePath, start, end, logsAnalyzerUtil)
		}(start, end)
	}

	wg.Wait()
	fmt.Printf("Количество логов: %d, Средний размер ответа: %d\n",
		logsAnalyzerUtil.LogsNumber, logsAnalyzerUtil.ResponseSizeSum/logsAnalyzerUtil.LogsNumber)
}

// Обрабатывает заданный диапазон байтов файла
func processChunk(filePath string, start, end int64, logAnalyzerUtil *LogAnalyzerUtil) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Ошибка открытия файла в горутине:", err)
		return
	}
	defer file.Close()

	// Переходим к началу чанка
	file.Seek(start, 0)

	// Читаем на несколько байтов больше для захвата конца строки
	bufferSize := end - start + 1024 // дополнительный буфер
	buffer := make([]byte, bufferSize)
	bytesRead, err := file.Read(buffer)
	if err != nil {
		fmt.Println("Ошибка чтения чанка:", err)
		return
	}
	buffer = buffer[:bytesRead]

	// Найти последний полный конец строки в буфере
	lastNewlineIndex := int64(-1)
	for i := len(buffer) - 1; i >= 0; i-- {
		if buffer[i] == '\n' {
			lastNewlineIndex = int64(i)
			break
		}
	}

	// Если найден конец строки, ограничиваем буфер до последней полной строки
	if lastNewlineIndex != -1 {
		buffer = buffer[:lastNewlineIndex+1]
	}

	// Разделяем прочитанный чанк на строки и обрабатываем каждую строку
	lines := regexp.MustCompile(`\r?\n`).Split(string(buffer), -1)
	for _, line := range lines {
		Parse(string(line), logAnalyzerUtil)
	}
}

func Parse(line string, logAnalyzerUtil *LogAnalyzerUtil) {
	re := regexp.MustCompile(`\[(\d{2}/\w{3}/\d{4}):.*?\] "(?:GET|HEAD|POST|PATCH) (/downloads/[\w\d_]+) HTTP/.*?" (\d{3}) (\d+)`)
	matches := re.FindStringSubmatch(line)

	// Проверка, что регулярное выражение нашло все необходимые группы
	if len(matches) < 5 {
		return
	}

	// Преобразование значений в int
	statusCode, _ := strconv.Atoi(matches[3])
	responseSize, _ := strconv.Atoi(matches[4])

	// Атомарные операции для общих переменных
	atomic.AddInt64(&logAnalyzerUtil.ResponseSizeSum, int64(responseSize))
	atomic.AddInt64(&logAnalyzerUtil.LogsNumber, 1)

	// Потокобезопасный доступ к картам с помощью мьютекса
	logAnalyzerUtil.mu.Lock()
	logAnalyzerUtil.MostRequestableResources[matches[2]]++
	logAnalyzerUtil.mostFrequentStatusCodes[int64(statusCode)]++
	logAnalyzerUtil.mu.Unlock()
}
