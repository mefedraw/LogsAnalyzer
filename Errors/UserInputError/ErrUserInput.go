package UserInputError

import "fmt"

type ErrUserInput struct {
	msg string
}

func NewErrUserInput(msg string) *ErrUserInput {
	return &ErrUserInput{msg: msg}
}

func (e *ErrUserInput) Error() string {
	return fmt.Sprintf("input error: %s", e.msg)
}
