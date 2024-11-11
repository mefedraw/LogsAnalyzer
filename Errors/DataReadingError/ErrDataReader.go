package DataReadingError

import "fmt"

type ErrDataReader struct {
	msg string
}

func NewErrDataReader(msg string) *ErrDataReader {
	return &ErrDataReader{msg: msg}
}

func (e *ErrDataReader) Error() string {
	return fmt.Sprintf("DataReader error: %s", e.msg)
}
