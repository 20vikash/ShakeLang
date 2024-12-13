package lexer

type Token struct {
	id     int
	Type_  Type
	Lexeme string
	line   int
	column int
}

func createToken(id int, type_ Type, lexeme string, line int, column int) Token {
	return Token{
		id:     id,
		Type_:  type_,
		Lexeme: lexeme,
		line:   line,
		column: column,
	}
}
