package BufferedSource

import (
	"NginxLogsAnalyzer/Errors/BufferedSourceError"
	"bufio"
	"net/http"
)

type HttpResponseReaderProvider struct{}

func NewHttpResponseReaderProvider() *HttpResponseReaderProvider {
	return &HttpResponseReaderProvider{}
}

func (rp *HttpResponseReaderProvider) DataBufferWrap(path string) (*bufio.Reader, error) {
	client := http.Client{}
	resp, err := client.Get(path)
	if err != nil {
		return nil, BufferedSourceError.NewErrHttpResponseReaderProvider(err.Error())
	}
	file := bufio.NewReader(resp.Body)
	return file, nil
}
