package lexer

import (
	"fmt"
	"os"
)

func createError(s string) {
	fmt.Fprintln(os.Stderr, "\033[31m", s, "\033[0m")
	os.Exit(1)
}

func quotesError() {
	createError("Thy string is a tragedy, bereft of its rightful bounds! Pray, ensure every opening hath a closing, else chaos shall reign.")
}

func decimalError() {
	createError("Lo, thy number is most unseemly, marred by an excess of dots! A single point sufficeth to part the whole from the fraction.")
}

func identifierError() {
	createError("Thy chosen name for thy variable or function is most unworthy! Pray, adhere to the sacred rules of naming, lest chaos befall thy code.")
}

func invalidLiteral() {
	createError("Thy literal is a jest most foul! Such a creation hath no place in this kingdom of logic. Pray, craft it anew with proper form.")
}
