package lexer

import (
	"fmt"
	"unicode"
)

func Lexer(s string) {
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
			numberEncountered = true
			temp += string(v)
		} else if numberEncountered && v == '.' {
			if decimalPoints == 1 {
				break
			}

			decimalPoints += 1
			temp += string(v)
		} else if v == '"' {
			if numberEncountered {
				break
			}
			if quoteEncountered {
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

			if v == ';' && !quoteEncountered {
				column = 0
				line += 1
			}
		}
	}

	fmt.Print(Tokens)
}
