package lexer

import (
	"unicode"
)

func checkValidVariableName(s string) bool {
	validSym := '_'
	if unicode.IsDigit(rune(s[0])) {
		return false
	}

	for _, v := range s {
		if !unicode.IsLetter(v) && v != validSym && !unicode.IsDigit(v) {
			return false
		}
	}

	return true
}

func addTemp(temp string, Tokens []Token, id int, line int, column int) ([]Token, int) {
	if len(temp) > 0 {
		r, exists := getType(temp)
		if exists {
			id++
			token := createToken(id, r, temp, line, column)
			Tokens = append(Tokens, token)
		} else {
			valid := checkValidVariableName(temp)
			if valid {
				id++
				token := createToken(id, IDENTIFIER, temp, line, column)
				Tokens = append(Tokens, token)
			} else {
				id++
				token := createToken(id, LITERAL, temp, line, column)
				Tokens = append(Tokens, token)
			}
		}
	}

	return Tokens, id
}

func Lexer(s string) []Token {
	s += "\n"

	Tokens := make([]Token, 0, 100)

	id := 0
	temp := ""
	line := 1
	column := 0

	quoteEncountered := false
	numberEncountered := false
	decimalPoints := 0

	for _, v := range s {
		if quoteEncountered && v != '"' {
			temp += string(v)
			continue
		}

		if unicode.IsDigit(v) {
			if len(temp) > 0 && !numberEncountered {
				id += 1
				v, exists := getType(temp)

				if exists {
					token := createToken(id, v, temp, line, column)
					Tokens = append(Tokens, token)
					temp = ""
				}
			}

			if len(temp) > 0 && !unicode.IsLetter(rune(temp[0])) {
				numberEncountered = true
			}
			temp += string(v)
		} else if numberEncountered && v == '.' {
			if decimalPoints == 1 {
				decimalError()
			}

			decimalPoints += 1
			temp += string(v)
		} else if v == '"' {
			if len(temp) > 0 && !quoteEncountered {
				id += 1
				v, exists := getType(temp)

				if exists {
					token := createToken(id, v, temp, line, column)
					Tokens = append(Tokens, token)
					temp = ""
				}
			}
			if numberEncountered {
				invalidLiteral()
			}
			if quoteEncountered {
				temp += string(v)
				id += 1
				token := createToken(id, LITERAL, temp, line, column)
				Tokens = append(Tokens, token)
				temp = ""
				quoteEncountered = false
				continue
			}
			quoteEncountered = true
			temp += string(v)
		} else if v != '\n' && v != ' ' {
			if numberEncountered {
				id += 1
				token := createToken(id, LITERAL, temp, line, column)
				Tokens = append(Tokens, token)
				temp = ""
				decimalPoints = 0
				numberEncountered = false
			} else {
				d, exists := getType(string(v))
				if exists {
					if quoteEncountered {
						quotesError()
					}
					id += 1
					Tokens, id = addTemp(temp, Tokens, id, line, column)

					token := createToken(id, d, string(v), line, column)
					Tokens = append(Tokens, token)
					temp = ""
					decimalPoints = 0
					continue
				}
			}

			d, exists := getType(temp)

			if exists {
				id++
				token := createToken(id, d, temp, line, column)
				Tokens = append(Tokens, token)
				temp = ""
			}

			temp += string(v)
			column += 1

		} else {
			if quoteEncountered {
				temp += string(v)
				continue
			}

			v, exists := getType(temp)

			if exists {
				if quoteEncountered {
					quotesError()
				}
				id += 1
				token := createToken(id, v, temp, line, column)
				Tokens = append(Tokens, token)
				temp = ""
			} else {
				if len(temp) > 0 {
					valid := checkValidVariableName(temp)

					if valid {
						id += 1
						token := createToken(id, IDENTIFIER, temp, line, column)
						Tokens = append(Tokens, token)
						temp = ""
					}
				}
			}
		}

		if v == '\n' {
			Tokens, id = addTemp(temp, Tokens, id, line, column)

			id++
			token := createToken(id, EOL, "EOL", line, column)
			Tokens = append(Tokens, token)

			column = 0
			line += 1
		}
	}

	// for _, v := range Tokens {
	// 	fmt.Println(v)
	// }

	return Tokens
}
