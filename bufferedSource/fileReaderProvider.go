package bufferedSource

import (
	"NginxLogsAnalyzer/Errors/BufferedSourceError"
	"bufio"
	"os"
)

type FileReaderProvider struct{}

func NewFileReaderProvider() *FileReaderProvider {
	return new(FileReaderProvider)
}

func (frp *FileReaderProvider) DataBufferWrap(path string) (*bufio.Reader, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, BufferedSourceError.NewErrorFileReaderProvider("file is not found")
	}

	file, err := os.Open(path)
	if err != nil {
		return nil, BufferedSourceError.NewErrorFileReaderProvider("is not possible to open file")
	}

	reader := bufio.NewReader(file)
	return reader, nil
}
