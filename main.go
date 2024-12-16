package main

import (
	"fmt"
	"shake/shake/backend"
	"shake/shake/frontend/lexer"
	"shake/shake/frontend/parser"
)

func main() {

	// if len(os.Args) < 2 {
	// 	fmt.Fprintln(os.Stderr, "\033[31mPray, provide the file path as an argument!\033[0m")
	// 	os.Exit(1)
	// }

	// fileName := os.Args[1]

	// if len(fileName) < 4 || fileName[len(fileName)-4:] != ".thy" {
	// 	fmt.Fprintln(os.Stderr, "\033[31mPray, furnish the proper path to thine file, and ensure its noble extension doth bear the mark of `.thy`.\033[0m")
	// 	os.Exit(1)
	// }

	// d, err := os.ReadFile("some.thy")

	// if err == nil {
	// 	tokens := lexer.Lexer(string(d))
	// 	parser.Ast(tokens)
	// }

	eq := "(1+10/3*(8+2)-1)"
	tokens := lexer.Lexer(string(eq))
	binex := parser.BinaryExpressionTree(tokens)
	variables := make(map[string]string, 0)

	fmt.Println(backend.EvaluateBinaryExpression(binex.Left, binex.Right, binex.Operator, variables))
}
