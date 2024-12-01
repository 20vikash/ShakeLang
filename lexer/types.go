package lexer

type Types int

const (
	KEYWORD Types = iota
	IDENTIFIER
	LITERAL
	OPERATOR
	PUNCTUATION
)
