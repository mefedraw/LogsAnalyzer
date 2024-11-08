package DataReading

import (
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
				return err
			}
		}
		lines <- line
	}
	return nil
}
