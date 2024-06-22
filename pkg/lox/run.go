package lox

import (
	"fmt"
)

type Lox struct {
}

func NewLox() *Lox {
	return &Lox{}
}

func (l *Lox) Run(source string) []Error {
	scanner := NewScanner(source)
	tokens, errors := scanner.ScanTokens()
	if len(errors) > 0 {
		return errors
	}
	for _, token := range tokens {
		fmt.Println(token)
	}
	return nil
}
