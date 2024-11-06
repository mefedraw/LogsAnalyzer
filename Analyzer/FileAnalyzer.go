package Analyzer

import (
	"NginxLogsAnalyzer/LogsUtil"
	"NginxLogsAnalyzer/Parsing"
	"fmt"
	"os"
	"regexp"
	"sync"
)

type FileAnalyzer struct{}

func NewFileAnalyzer() *FileAnalyzer {
	return &FileAnalyzer{}
}

func (fa *FileAnalyzer) Analyze(path string, parser Parsing.LogsParser) *LogsUtil.LogAnalyzerUtil {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Ошибка открытия файла:", err)
		return nil
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("Ошибка получения информации о файле:", err)
		return nil
	}

	fileSize := fileInfo.Size()
	numWorkers := 16 // Количество горутин
	chunkSize := fileSize / int64(numWorkers)

	logsAnalyzerUtil := LogsUtil.NewLogAnalyzerUtil()
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
			fa.processChunk(path, start, end, logsAnalyzerUtil)
		}(start, end)
	}

	wg.Wait()
	fmt.Printf("Количество логов: %d, Средний размер ответа: %d\n",
		logsAnalyzerUtil.LogsNumber, logsAnalyzerUtil.ResponseSizeSum/logsAnalyzerUtil.LogsNumber)

	return logsAnalyzerUtil
}

func (fa *FileAnalyzer) processChunk(filePath string, start, end int64, logAnalyzerUtil *LogsUtil.LogAnalyzerUtil) {
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

func (fa *FileAnalyzer) findLineEndingSymbol(buffer []byte) int64 {
	for i := len(buffer) - 1; i >= 0; i-- {
		if buffer[i] == '\n' {
			return int64(i)
		}
	}

	return -1
}
