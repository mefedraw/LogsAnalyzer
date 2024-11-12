package bufferedSource

import (
	"NginxLogsAnalyzer/Errors/BufferedSourceError"
	"bufio"
	"fmt"
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

	if resp.StatusCode != http.StatusOK {
		return nil, BufferedSourceError.NewErrHttpResponseReaderProvider(
			fmt.Sprintf("unexpected status code: %d", resp.StatusCode),
		)
	}

	file := bufio.NewReader(resp.Body)
	return file, nil
}
