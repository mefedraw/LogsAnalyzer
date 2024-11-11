package BufferedSourceError

import "fmt"

type ErrHttpResponseReaderProvider struct {
	msg string
}

func NewErrHttpResponseReaderProvider(msg string) *ErrHttpResponseReaderProvider {
	return &ErrHttpResponseReaderProvider{msg: msg}
}

func (e *ErrHttpResponseReaderProvider) Error() string {
	return fmt.Sprintf("HttpResponseReaderProvider error: %s", e.msg)
}
