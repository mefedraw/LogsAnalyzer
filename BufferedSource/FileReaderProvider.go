package BufferedSource

import (
	"bufio"
	"os"
)

type FileReaderProvider struct{}

func NewFileReaderProvider() *FileReaderProvider {
	return new(FileReaderProvider)
}

func (frp *FileReaderProvider) DataBufferWrap(path string) (*bufio.Reader, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// defer file.Close()
	return bufio.NewReader(file), nil
}
