package BufferedSourceError

import "fmt"

type ErrorFileReaderProvider struct {
	msg string
}

func NewErrorFileReaderProvider(msg string) *ErrorFileReaderProvider {
	return &ErrorFileReaderProvider{msg: msg}
}

func (e *ErrorFileReaderProvider) Error() string {
	return fmt.Sprintf("FileReaderProvider error: %s", e.msg)
}
