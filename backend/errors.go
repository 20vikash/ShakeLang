package backend

import (
	"fmt"
	"os"
)

func createError(s string) {
	fmt.Fprintln(os.Stderr, "\033[31m", s, "\033[0m")
	os.Exit(1)
}

func variableNotDeclared() {
	createError("A phantom name thou hast invoked, yet no such variable was ever declared in thy realm!")
}

func variableAlreadyExists() {
	createError("Thy attempt to declare is in vain, for the variable already walks this realm!")
}

func invalidBinaryExpressionValue() {
	createError("Arithmetic defies thy will, for strings have no place among numbers!")
}
