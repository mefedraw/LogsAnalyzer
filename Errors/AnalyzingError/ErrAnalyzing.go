﻿package AnalyzingError

import "fmt"

type ErrAnalyzing struct {
	msg string
}

func NewErrAnalyzing(msg string) *ErrAnalyzing {
	return &ErrAnalyzing{msg: msg}
}

func (e *ErrAnalyzing) Error() string {
	return fmt.Sprintf("analyzing error: %s", e.msg)
}
