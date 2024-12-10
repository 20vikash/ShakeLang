package parser

import (
	"fmt"
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

func Postfix(infix string) {
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
				}

				postfix += temp
			}
		} else {
			if len(stack) == 0 || (priority[string(v)] > priority[stack[len(stack)-1]]) {
				stack = append(stack, string(v))
			} else {
				for priority[string(v)] <= priority[stack[len(stack)-1]] {
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

	fmt.Println(postfix)
}
