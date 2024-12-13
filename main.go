package main

import (
	"os"
	"shake/shake/lexer"
	"shake/shake/parser"
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

	d, err := os.ReadFile("some.thy")

	if err == nil {
		tokens := lexer.Lexer(string(d))
		parser.Ast(tokens)
	}
}
