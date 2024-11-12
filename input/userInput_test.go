package input_test

import (
	input2 "NginxLogsAnalyzer/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInput_ParseInput_IncorrectCommand_Failure(t *testing.T) {
	input := "analyze --path nginx_logs.txt"

	inputData := input2.NewUserInput()
	err := inputData.ParseInput(input)

	assert.Error(t, err)
	assert.EqualError(t, err, "input error: incorrect command")
}

func TestInput_ParseInput_DefaultCommand_Success(t *testing.T) {
	input := "analyzer --path nginx_logs.txt"

	inputData := input2.NewUserInput()
	err := inputData.ParseInput(input)

	assert.Nil(t, err)
	assert.Equal(t, inputData.Path, "nginx_logs.txt")
	assert.Equal(t, inputData.Format, "markdown")
}

func TestInput_ParseInput_CommandWithFormat_Success(t *testing.T) {
	input := "analyzer --path nginx_logs.txt --format adoc"

	inputData := input2.NewUserInput()
	err := inputData.ParseInput(input)

	assert.Nil(t, err)
	assert.Equal(t, inputData.Format, "adoc")
}

func TestInput_ParseInput_CommandWithToDate_Success(t *testing.T) {
	input := "analyzer --path nginx_logs.txt --to 2015-05-17"

	inputData := input2.NewUserInput()
	err := inputData.ParseInput(input)

	assert.Nil(t, err)
	assert.Equal(t, inputData.ToDate, "2015-05-17")
}

func TestInput_ParseInput_CommandWithFromDate_Success(t *testing.T) {
	input := "analyzer --path nginx_logs.txt --from 2015-05-18"

	inputData := input2.NewUserInput()
	err := inputData.ParseInput(input)

	assert.Nil(t, err)
	assert.Equal(t, inputData.FromDate, "2015-05-18")
}

func TestInput_ParseInput_UrlPath_Success(t *testing.T) {
	input := "analyzer --path https://raw.githubusercontent.com/elastic/examples/master/Common%20Data%20Formats/nginx_logs/nginx_logs"

	inputData := input2.NewUserInput()
	err := inputData.ParseInput(input)

	assert.Nil(t, err)
	assert.Equal(t, inputData.Path, "https://raw.githubusercontent.com/elastic/examples/master/Common%20Data%20Formats/nginx_logs/nginx_logs")
}
