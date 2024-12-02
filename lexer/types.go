package lexer

type Type int

const (
	KEYWORD Type = iota
	IDENTIFIER
	LITERAL
	OPERATOR
	PUNCTUATOR
	nil
)

func getType(s string) (Type, bool) {
	types := map[string]Type{
		"declare":   KEYWORD,
		"sonnet":    KEYWORD,
		"prechance": KEYWORD,
		"yield":     KEYWORD,
		"elsewise":  KEYWORD,
		"proclaim":  KEYWORD,
		"forsooth":  KEYWORD,
		"whilst":    KEYWORD,
		"but":       OPERATOR,
		"an'":       OPERATOR,
		"or'":       OPERATOR,
		"giveth":    OPERATOR,
		"+=":        OPERATOR,
		"-=":        OPERATOR,
		"+":         OPERATOR,
		"-":         OPERATOR,
		"*":         OPERATOR,
		"/":         OPERATOR,
		"*=":        OPERATOR,
		"/=":        OPERATOR,
		"%":         OPERATOR,
		"%=":        OPERATOR,
		",":         PUNCTUATOR,
		"{":         PUNCTUATOR,
		"}":         PUNCTUATOR,
		"(":         PUNCTUATOR,
		")":         PUNCTUATOR,
		"[":         PUNCTUATOR,
		"]":         PUNCTUATOR,
		";":         PUNCTUATOR,
	}

	v, exists := types[s]

	if exists {
		return v, true
	}

	return nil, false
}
