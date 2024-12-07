package lexer

import (
	"fmt"
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

func Lexer(s string) {
	s += "\n~"

	Tokens := make([]Token, 0, 100)

	id := 0
	temp := ""
	line := 1
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
					if len(temp) > 0 {
						r, exists := getType(temp)
						if exists {
							token := createToken(id, r, temp, line, column)
							Tokens = append(Tokens, token)
						} else {
							valid := checkValidVariableName(temp)
							if valid {
								token := createToken(id, IDENTIFIER, temp, line, column)
								Tokens = append(Tokens, token)
							} else {
								identifierError()
							}
						}
					}
					token := createToken(id, d, string(v), line, column)
					Tokens = append(Tokens, token)
					temp = ""
					decimalPoints = 0
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
			if !quoteEncountered {
				quotesError()
			}

			column = 0
			line += 1
		}
	}

	for _, v := range Tokens {
		fmt.Println(v)
	}
}
