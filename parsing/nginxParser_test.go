package parsing_test

import (
	"NginxLogsAnalyzer/parsing"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNginxLogsParser_ParseLine_Success(t *testing.T) {
	parser := parsing.NewNginxLogsParser()
	line := `93.180.71.3 - - [17/May/2015:08:05:32 +0000] "GET /downloads/product_1 HTTP/1.1" 304 0`

	matches := parser.ParseLine(line)

	expected := []string{
		"93.180.71.3 - - [17/May/2015:08:05:32 +0000] \"GET /downloads/product_1 HTTP/1.1\" 304 0",
		"93.180.71.3",
		"17/May/2015",
		"/downloads/product_1",
		"304",
		"0",
	}
	assert.NotNil(t, matches)
	assert.Equal(t, &expected, matches)
}

func TestNginxLogsParser_ParseLine_IncompleteLine(t *testing.T) {
	parser := parsing.NewNginxLogsParser()
	line := `93.180.71.3 - - [17/May/2015:08:05:32 +0000] "GET /downloads/product_1 HTTP/1.1" 304`

	matches := parser.ParseLine(line)

	assert.Nil(t, matches)
}
