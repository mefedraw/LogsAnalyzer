package FileReaderProviderServiceError

import "fmt"

type ErrorFileReaderProvider struct {
	msg string
}

func NewErrorFileReaderProvider(msg string) *ErrorFileReaderProvider {
	return &ErrorFileReaderProvider{msg: msg}
}

func (erp *ErrorFileReaderProvider) Error() string {
	return fmt.Sprintf("NginxParser error: %s", erp.msg)
}
