package lox

import (
	"fmt"
)

type Error struct {
	Message string
	Line    int
}

func (e *Error) String() string {
	return fmt.Sprintf("[line %d] %s", e.Line, e.Message)
}

func (e *Error) Error() string {
	return e.String()
}

func NewError(line int, message string) *Error {
	return &Error{
		Message: message,
		Line:    line,
	}
}
