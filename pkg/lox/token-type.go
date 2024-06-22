package lox

type TokenType rune

const (
	// Multi-character tokens
	BANG_EQUAL TokenType = -(iota + 1)
	EQUAL_EQUAL
	GREATER_EQUAL
	LESS_EQUAL

	// Literals
	IDENTIFIER
	STRING
	NUMBER

	// Keywords
	AND
	CLASS
	ELSE
	FALSE
	FUNCTI
	FOR
	IF
	NIL
	OR
	PRINT
	RETURN
	SUPER
	THIS
	TRUE
	VAR
	WHILE
	EOF

	// Single-character tokens
	LEFT_PAREN  TokenType = '('
	RIGHT_PAREN TokenType = ')'
	LEFT_BRACE  TokenType = '{'
	RIGHT_BRACE TokenType = '}'
	COMMA       TokenType = ','
	DOT         TokenType = '.'
	MINUS       TokenType = '-'
	PLUS        TokenType = '+'
	SEMICOLON   TokenType = ';'
	SLASH       TokenType = '/'
	STAR        TokenType = '*'
	BANG        TokenType = '!'
	EQUAL       TokenType = '='
	LESS        TokenType = '<'
	GREATER     TokenType = '>'
)

// ChatGPT made this one, Im too lazy to do it myself
func (t TokenType) String() string {
	switch t {
	case BANG_EQUAL:
		return "BANG_EQUAL"
	case EQUAL_EQUAL:
		return "EQUAL_EQUAL"
	case GREATER_EQUAL:
		return "GREATER_EQUAL"
	case LESS_EQUAL:
		return "LESS_EQUAL"
	case IDENTIFIER:
		return "IDENTIFIER"
	case STRING:
		return "STRING"
	case NUMBER:
		return "NUMBER"
	case AND:
		return "AND"
	case CLASS:
		return "CLASS"
	case ELSE:
		return "ELSE"
	case FALSE:
		return "FALSE"
	case FUNCTI:
		return "FUNCTI"
	case FOR:
		return "FOR"
	case IF:
		return "IF"
	case NIL:
		return "NIL"
	case OR:
		return "OR"
	case PRINT:
		return "PRINT"
	case RETURN:
		return "RETURN"
	case SUPER:
		return "SUPER"
	case THIS:
		return "THIS"
	case TRUE:
		return "TRUE"
	case VAR:
		return "VAR"
	case WHILE:
		return "WHILE"
	case EOF:
		return "EOF"
	case LEFT_PAREN:
		return "LEFT_PAREN"
	case RIGHT_PAREN:
		return "RIGHT_PAREN"
	case LEFT_BRACE:
		return "LEFT_BRACE"
	case RIGHT_BRACE:
		return "RIGHT_BRACE"
	case COMMA:
		return "COMMA"
	case DOT:
		return "DOT"
	case MINUS:
		return "MINUS"
	case PLUS:
		return "PLUS"
	case SEMICOLON:
		return "SEMICOLON"
	case SLASH:
		return "SLASH"
	case STAR:
		return "STAR"
	case BANG:
		return "BANG"
	case EQUAL:
		return "EQUAL"
	case LESS:
		return "LESS"
	case GREATER:
		return "GREATER"
	default:
		return "UNKNOWN"
	}
}

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"functi": FUNCTI,
	"for":    FOR,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}
