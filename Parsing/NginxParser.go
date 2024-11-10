package Parsing

import (
	"fmt"
	"regexp"
	"strings"
)

type NginxLogsParser struct {
}

func NewNginxLogsParser() *NginxLogsParser {
	return &NginxLogsParser{}
}

func (nlp *NginxLogsParser) ParseLine(line string) *[]string {
	// Удаление BOM или других невидимых символов
	line = strings.TrimLeft(line, "\ufeff")

	re := regexp.MustCompile(`^([\da-fA-F:.]+) - - \[(\d{2}/\w{3}/\d{4}):\d{2}:\d{2}:\d{2} [+\-]\d{4}\] "\w+ ([^ ]+) HTTP/[\d.]+" (\d{3}) (\d+)`)
	matches := re.FindStringSubmatch(line)

	if len(matches) < 6 {
		fmt.Println(line)
		panic("пиздец")
	}
	return &matches
}
