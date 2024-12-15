package parser

import (
	"fmt"
	"reflect"
	"shake/shake/frontend/lexer"
	"unicode"
)

var priority = map[string]int{
	"^": 3,
	"/": 2,
	"*": 2,
	"+": 1,
	"-": 1,
	"(": 0,
}

func postfix(infix string) string {
	stack := make([]string, 0)
	postfix := ""
	chunk := ""

	for _, v := range infix {
		if v != ',' {
			chunk += string(v)
			continue
		}

		if chunk == "(" {
			stack = append(stack, string(chunk))
			chunk = ""
			continue
		}
		if unicode.IsDigit(rune(chunk[0])) || unicode.IsLetter(rune(chunk[0])) {
			postfix += fmt.Sprintf("%v,", chunk)
		} else if chunk == ")" {
			for {
				temp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if temp == "(" {
					break
				} else {
					postfix += fmt.Sprintf("%v,", temp)
				}
			}
		} else {
			if len(stack) == 0 || (priority[string(chunk)] > priority[stack[len(stack)-1]]) {
				stack = append(stack, string(chunk))
			} else {
				for len(stack) > 0 && priority[string(chunk)] <= priority[stack[len(stack)-1]] {
					temp := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					postfix += fmt.Sprintf("%v,", temp)
				}

				stack = append(stack, string(chunk))
			}
		}

		chunk = ""
	}

	if len(stack) > 0 {
		for i := len(stack) - 1; i >= 0; i-- {
			postfix += fmt.Sprintf("%v,", stack[i])
		}
	}

	return postfix
}

func binaryExpressionTree(tokens []lexer.Token) BinaryExpression {
	infix := ""

	for _, v := range tokens {
		if v.Lexeme == "EOL" {
			break
		}
		infix += fmt.Sprintf("%v,", v.Lexeme)
	}

	postfix := postfix(infix)
	binExpression := getBinaryExpression(postfix)
	// fmt.Println(binExpression.right)

	return binExpression
}

func getNodeTypeString(value any) Node {
	if unicode.IsLetter(rune(value.(string)[0])) || value.(string)[0] == '_' {
		return createIdentifierNode(value.(string))
	} else {
		return createLiteralNode(value.(string))
	}
}

func getBinaryExpression(postfix string) BinaryExpression {
	stack := make([]any, 0)
	chunk := ""

	for _, v := range postfix {
		if v != ',' {
			chunk += string(v)
			continue
		}

		if unicode.IsDigit(rune(chunk[0])) || unicode.IsLetter(rune(chunk[0])) {
			stack = append(stack, chunk)
		} else {
			if len(stack) < 2 {
				break //TODO: Invalid arithmetic expression
			}

			rightValue := stack[len(stack)-1]
			leftValue := stack[len(stack)-2]

			rightType := fmt.Sprintf("%v", reflect.TypeOf(rightValue))
			leftType := fmt.Sprintf("%v", reflect.TypeOf(leftValue))

			if rightType == "string" && leftType == "string" {
				bin := BinaryExpression{
					_type:    "BinaryExpression",
					left:     getNodeTypeString(leftValue),
					right:    getNodeTypeString(rightValue),
					operator: chunk,
				}

				stack = stack[:len(stack)-2]

				stack = append(stack, bin)
			} else if rightType == "parser.BinaryExpression" && leftType == "string" {
				bin := BinaryExpression{
					_type:    "BinaryExpression",
					left:     getNodeTypeString(leftValue),
					right:    stack[len(stack)-1].(BinaryExpression),
					operator: chunk,
				}

				stack = stack[:len(stack)-2]

				stack = append(stack, bin)
			} else if rightType == "string" && leftType == "parser.BinaryExpression" {
				bin := BinaryExpression{
					_type:    "BinaryExpression",
					left:     stack[len(stack)-2].(BinaryExpression),
					right:    getNodeTypeString(rightValue),
					operator: chunk,
				}

				stack = stack[:len(stack)-2]

				stack = append(stack, bin)
			} else {
				bin := BinaryExpression{
					_type:    "BinaryExpression",
					left:     stack[len(stack)-2].(BinaryExpression),
					right:    stack[len(stack)-1].(BinaryExpression),
					operator: chunk,
				}

				stack = stack[:len(stack)-2]

				stack = append(stack, bin)
			}
		}

		chunk = ""
	}

	return stack[0].(BinaryExpression)
}
