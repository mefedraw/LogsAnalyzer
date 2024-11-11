package DataCollectingError

import "fmt"

type ErrLogDataCollector struct {
	msg string
}

func NewErrLogDataCollector(msg string) *ErrLogDataCollector {
	return &ErrLogDataCollector{msg: msg}
}

func (el *ErrLogDataCollector) Error() string {
	return fmt.Sprintf("LogDataCollector error: %s", el.msg)
}
