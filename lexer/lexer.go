package lexer

import (
	"fmt"
	"unicode"
)

func Lexer(s string) {
	s += "\n~"

	Tokens := make([]Token, 0, 100)

	id := 0
	temp := ""
	line := 0
	column := 0

	quoteEncountered := false
	numberEncountered := false
	decimalPoints := 0

	for _, v := range s {
		if unicode.IsDigit(v) && !quoteEncountered {
			if len(temp) > 0 && !numberEncountered {
				id += 1
				v, exists := getType(temp)

				if exists {
					token := createToken(id, v, temp, line, column)
					Tokens = append(Tokens, token)
					temp = ""
				}
			}

			numberEncountered = true
			temp += string(v)
		} else if numberEncountered && v == '.' {
			if decimalPoints == 1 {
				break
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
				break
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
				numberEncountered = false
			} else {
				d, exists := getType(string(v))
				if exists {
					id += 1
					if len(temp) > 0 {
						token := createToken(id, d, temp, line, column)
						Tokens = append(Tokens, token)
					}
					token := createToken(id, d, string(v), line, column)
					Tokens = append(Tokens, token)
					temp = ""
					continue
				}
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
				id += 1
				token := createToken(id, v, temp, line, column)
				Tokens = append(Tokens, token)
				temp = ""
			}

			if v == '\n' && !quoteEncountered {
				column = 0
				line += 1
			}
		}
	}

	for _, v := range Tokens {
		fmt.Println(v)
	}
}
