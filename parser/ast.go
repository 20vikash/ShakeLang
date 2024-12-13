package parser

import (
	"fmt"
	"shake/shake/lexer"
)

func Ast(tokens []lexer.Token) {
	nodes := make([]Node, 0)
	count := 0

	temp := make([]lexer.Token, 0)

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Lexeme == "declare" {
			if count == 0 && tokens[i+1].Type_ == lexer.IDENTIFIER && tokens[i+2].Type_ == lexer.EOL {
				idNode := createIdentifierNode(tokens[i+1].Lexeme)
				declareNode := createVariableDeclarationNodeWithoutInit(idNode)
				nodes = append(nodes, declareNode)
			} else if count == 0 && tokens[i+1].Type_ == lexer.IDENTIFIER && tokens[i+2].Lexeme == "giveth" && tokens[i+3].Type_ != lexer.EOL {
				id := createIdentifierNode(tokens[i+1].Lexeme)

				if tokens[i+4].Type_ != lexer.EOL {
					num := i + 3

					for tokens[num].Lexeme != "EOL" {
						temp = append(temp, tokens[num])
						num++
					}

					node := binaryExpressionTree(temp)
					declareNode := createVariableDeclarationNode(id, node)
					nodes = append(nodes, declareNode)
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
						break //TODO: Invalid init value
					}
				}
			} else {
				break //TODO: Invalid statement
			}
		}

		if tokens[i].Lexeme == "proclaim" {
			if count == 0 && tokens[i+1].Type_ != lexer.EOL {
				if tokens[i+2].Type_ != lexer.EOL {
					num := i + 1

					for tokens[num].Lexeme != "EOL" {
						temp = append(temp, tokens[num])
						num++
					}

					node := binaryExpressionTree(temp)
					proclaimNode := createProclaimStatementNode(node)
					nodes = append(nodes, proclaimNode)
				} else {
					if tokens[i+1].Type_ == lexer.LITERAL {
						literalNode := createLiteralNode(tokens[i+1].Lexeme)
						proclaimNode := createProclaimStatementNode(literalNode)
						nodes = append(nodes, proclaimNode)
					} else if tokens[i+1].Type_ == lexer.IDENTIFIER {
						idNode := createIdentifierNode(tokens[i+1].Lexeme)
						proclaimNode := createProclaimStatementNode(idNode)
						nodes = append(nodes, proclaimNode)
					} else {
						break //TODO: Invalid arguement.
					}
				}
			} else {
				break //TODO: Invalid proclaim statement
			}
		}

		if tokens[i].Lexeme != "EOL" {
			count++
			continue
		}

		temp = temp[:0]
		count = 0
	}

	programNode := createProgramNode(nodes)
	fmt.Println(programNode)
}
