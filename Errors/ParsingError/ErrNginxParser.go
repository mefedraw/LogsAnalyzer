package ParsingError

import "fmt"

type ErrNginxParser struct {
	msg string
}

func NewErrNginxParser(msg string) *ErrNginxParser {
	return &ErrNginxParser{msg: msg}
}

func (e *ErrNginxParser) Error() string {
	return fmt.Sprintf("NginxParser error: %s", e.msg)
}
