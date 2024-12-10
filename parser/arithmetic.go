package parser

import (
	"fmt"
	"reflect"
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

	for _, v := range infix {
		if v == '(' {
			stack = append(stack, string(v))
			continue
		}
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			postfix += string(v)
		} else if v == ')' {
			for {
				temp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if temp == "(" {
					break
				} else {
					postfix += temp
				}
			}
		} else {
			if len(stack) == 0 || (priority[string(v)] > priority[stack[len(stack)-1]]) {
				stack = append(stack, string(v))
			} else {
				for len(stack) > 0 && priority[string(v)] <= priority[stack[len(stack)-1]] {
					temp := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					postfix += temp
				}

				stack = append(stack, string(v))
			}
		}
	}

	if len(stack) > 0 {
		for i := len(stack) - 1; i >= 0; i-- {
			postfix += stack[i]
		}
	}

	return postfix
}

func GetBinaryExpression(infix string) {
	postfix := postfix(infix)
	stack := make([]any, 0)

	for _, v := range postfix {
		if unicode.IsDigit(v) || unicode.IsLetter(v) {
			stack = append(stack, string(v))
		} else {
			rightValue := stack[len(stack)-1]
			leftValue := stack[len(stack)-2]

			rightType := fmt.Sprintf("%v", reflect.TypeOf(rightValue))
			leftType := fmt.Sprintf("%v", reflect.TypeOf(leftValue))

			if rightType == "string" && leftType == "string" {
				bin := BinaryExpression{
					_type:    "BinaryExpression",
					left:     Literal{_type: "Literal", value: leftValue.(string)},
					right:    Literal{_type: "Literal", value: rightValue.(string)},
					operator: string(v),
				}

				stack = stack[:len(stack)-2]

				stack = append(stack, bin)
			} else if rightType == "parser.BinaryExpression" && leftType == "string" {
				bin := BinaryExpression{
					_type:    "BinaryExpression",
					left:     Literal{_type: "Literal", value: leftValue.(string)},
					right:    stack[len(stack)-1].(BinaryExpression),
					operator: string(v),
				}

				stack = stack[:len(stack)-2]

				stack = append(stack, bin)
			} else if rightType == "string" && leftType == "parser.BinaryExpression" {
				bin := BinaryExpression{
					_type:    "BinaryExpression",
					left:     stack[len(stack)-2].(BinaryExpression),
					right:    Literal{_type: "Literal", value: rightValue.(string)},
					operator: string(v),
				}

				stack = stack[:len(stack)-2]

				stack = append(stack, bin)
			} else {
				bin := BinaryExpression{
					_type:    "BinaryExpression",
					left:     stack[len(stack)-2].(BinaryExpression),
					right:    stack[len(stack)-1].(BinaryExpression),
					operator: string(v),
				}

				stack = stack[:len(stack)-2]

				stack = append(stack, bin)
			}
		}
	}

	fmt.Println(stack[0].(BinaryExpression).right)
}
