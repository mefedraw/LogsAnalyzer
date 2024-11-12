package dataReading

import (
	"bufio"
)

type DataReader interface {
	ReadBuffer(reader *bufio.Reader, lines chan string) error
}
