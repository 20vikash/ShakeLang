package parser

import (
	"fmt"
	"os"
)

func createError(s string) {
	fmt.Fprintln(os.Stderr, "\033[31m", s, "\033[0m")
	os.Exit(1)
}

func invalidInitValue() {
	createError("An ill-formed beginning, thy value is unfit for initialization. Amend thy ways!")
}

func invalidDeclaration() {
	createError("A flawed assertion, thy declaration is unworthy. Rewrite with wisdom!")
}

func invalidProclamation() {
	createError("A flawed assertion, thy proclamation is unworthy. Rewrite with wisdom!")
}

func invalidAssignment() {
	createError("A flawed assertion, thy assignment is unworthy. Rewrite with wisdom!")
}

func invalidArgProclaim() {
	createError("An ill-formed beginning, thy value is unfit for proclamation. Amend thy ways!")
}

func invalidStatement() {
	createError("A wayward utterance, thy statement strays from reason. Speak with clarity anew!")
}
