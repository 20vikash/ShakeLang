package lexer

type Token struct {
	id     int
	type_  Type
	lexeme string
	line   int
	column int
}

func createToken(id int, type_ Type, lexeme string, line int, column int) Token {
	return Token{
		id:     id,
		type_:  type_,
		lexeme: lexeme,
		line:   line,
		column: column,
	}
}
