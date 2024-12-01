package lexer

type token struct {
	id     int
	type_  Types
	lexeme string
	line   int
	column int
}

func (t token) CreateToken(id int, type_ Types, lexeme string, line int, column int) *token {
	return &token{
		id:     id,
		type_:  type_,
		lexeme: lexeme,
		line:   line,
		column: column,
	}
}
