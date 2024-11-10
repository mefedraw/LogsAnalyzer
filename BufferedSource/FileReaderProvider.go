package BufferedSource

import (
	"bufio"
	"fmt"
	"os"
)

type FileReaderProvider struct{}

func NewFileReaderProvider() *FileReaderProvider {
	return new(FileReaderProvider)
}

func (frp *FileReaderProvider) DataBufferWrap(path string) (*bufio.Reader, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("файл не найден: %s", path)
	}

	// Открываем файл
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("ошибка при открытии файла %s: %v", path, err)
	}

	reader := bufio.NewReader(file)
	return reader, nil
}
