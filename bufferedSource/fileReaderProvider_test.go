package bufferedSource_test

import (
	"NginxLogsAnalyzer/bufferedSource"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
)

func TestBufferedSource_DataBufferWrap_FileExists_Success(t *testing.T) {
	tmpFile, err := ioutil.TempFile("", "testfile")
	assert.NoError(t, err)
	defer os.Remove(tmpFile.Name())

	content := "test data"
	_, err = tmpFile.WriteString(content)
	assert.NoError(t, err)
	err = tmpFile.Close()
	assert.NoError(t, err)

	frp := bufferedSource.NewFileReaderProvider()
	reader, err := frp.DataBufferWrap(tmpFile.Name())
	assert.NoError(t, err)
	assert.NotNil(t, reader)
}

func TestBufferedSource_DataBufferWrap_FileNotFound_Error(t *testing.T) {
	frp := bufferedSource.NewFileReaderProvider()
	_, err := frp.DataBufferWrap("nonexistentfile.txt")

	assert.Error(t, err)
	assert.EqualError(t, err, "FileReaderProvider error: file is not found")
}
