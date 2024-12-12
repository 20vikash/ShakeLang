package parser

import (
	"shake/shake/lexer"
)

func Ast(tokens []lexer.Token) {
	nodes := make([]Node, 0)
	count := 0

	for i, v := range tokens {
		if v.Lexeme == "declare" {
			if count == 0 && tokens[i+1].Type_ == lexer.IDENTIFIER && tokens[i+2].Lexeme == "giveth" && tokens[i+3].Type_ != lexer.EOL {
				id := createIdentifierNode(tokens[i+1].Lexeme)

				if tokens[i+4].Type_ != lexer.EOL {

				} else {
					if tokens[i+3].Type_ == lexer.LITERAL {
						literalNode := createLiteralNode(tokens[i+3].Lexeme)
						declareNode := createVariableDeclarationNode(id, literalNode)
						nodes = append(nodes, declareNode)
					} else if tokens[i+3].Type_ == lexer.IDENTIFIER {
						idNode := createIdentifierNode(tokens[i+3].Lexeme)
						declareNode := createVariableDeclarationNode(id, idNode)
						nodes = append(nodes, declareNode)
					} else {
						break
					}
				}
			} else {
				break
			}
		}
	}
}
