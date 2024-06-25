package lox

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

	parser := NewParser(tokens)
	expr := parser.Parse()
	PrintAST(expr)
	return nil
}
