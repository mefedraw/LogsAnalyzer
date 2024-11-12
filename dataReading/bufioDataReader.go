package dataReading

import (
	"NginxLogsAnalyzer/Errors/DataReadingError"
	"bufio"
	"io"
)

type BufioDataReader struct{}

func NewBufioDataReader() *BufioDataReader {
	return &BufioDataReader{}
}

func (bd *BufioDataReader) ReadBuffer(reader *bufio.Reader, lines chan string) error {
	defer close(lines)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return DataReadingError.NewErrDataReader(err.Error())
			}
		}
		lines <- line
	}
	return nil
}
