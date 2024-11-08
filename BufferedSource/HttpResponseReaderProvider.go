package BufferedSource

import (
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
	defer resp.Body.Close()
	if err != nil {
		return nil, err
	}
	file := bufio.NewReader(resp.Body)
	return file, nil
}
