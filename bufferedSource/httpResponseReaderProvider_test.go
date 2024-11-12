package bufferedSource_test

import (
	"NginxLogsAnalyzer/bufferedSource"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpResponseReaderProvider_DataBufferWrap_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	rp := bufferedSource.NewHttpResponseReaderProvider()
	reader, err := rp.DataBufferWrap(server.URL)

	assert.NoError(t, err)
	assert.NotNil(t, reader)
}

func TestHttpResponseReaderProvider_DataBufferWrap_Error(t *testing.T) {
	rp := bufferedSource.NewHttpResponseReaderProvider()
	_, err := rp.DataBufferWrap("https://trolling_link")

	assert.Error(t, err)
	assert.EqualError(t, err, fmt.Sprintf("%s", err.Error()))
}
