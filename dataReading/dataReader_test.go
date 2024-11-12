package dataReading_test

import (
	"NginxLogsAnalyzer/dataReading"
	"bufio"
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBufioDataReader_ReadBuffer_Success(t *testing.T) {
	data := "line1\nline2\nline3\n"
	reader := bufio.NewReader(bytes.NewBufferString(data))
	lines := make(chan string)
	dataReader := dataReading.NewBufioDataReader()

	go func() {
		err := dataReader.ReadBuffer(reader, lines)
		assert.NoError(t, err)
	}()

	var result []string
	for line := range lines {
		result = append(result, line)
	}

	expected := []string{"line1\n", "line2\n", "line3\n"}
	assert.Equal(t, expected, result)
}
