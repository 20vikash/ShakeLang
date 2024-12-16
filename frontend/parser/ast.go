package parser

import (
	"fmt"
	"shake/shake/frontend/lexer"
)

func Ast(tokens []lexer.Token) {
	nodes := make([]Node, 0)
	count := 0
	inStatement := false

	temp := make([]lexer.Token, 0)

	for i := 0; i < len(tokens); i++ {
		if tokens[i].Type_ == lexer.EOL {
			temp = temp[:0]
			count = 0
			inStatement = false
			continue
		}

		if tokens[i].Lexeme == "declare" && !inStatement {
			inStatement = true
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

					node := BinaryExpressionTree(temp)
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
						invalidInitValue()
					}
				}
			} else {
				invalidDeclaration()
			}

			count++
		} else if tokens[i].Lexeme == "proclaim" && !inStatement {
			inStatement = true

			if count == 0 && tokens[i+1].Type_ != lexer.EOL {
				if tokens[i+2].Type_ != lexer.EOL {
					num := i + 1

					for tokens[num].Lexeme != "EOL" {
						temp = append(temp, tokens[num])
						num++
					}

					node := BinaryExpressionTree(temp)
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
						invalidArgProclaim()
					}
				}
			} else {
				invalidProclamation()
			}

			count++
		} else if tokens[i].Type_ == lexer.IDENTIFIER && !inStatement {
			inStatement = true

			if count == 0 && tokens[i+1].Lexeme == "giveth" && tokens[i+2].Type_ != lexer.EOL {
				id := createIdentifierNode(tokens[i].Lexeme)

				if tokens[i+3].Type_ != lexer.EOL {
					num := i + 2

					for tokens[num].Type_ != lexer.EOL {
						temp = append(temp, tokens[num])
						num++
					}

					node := BinaryExpressionTree(temp)
					initNode := createInitializationExpressionNode(id, node)
					nodes = append(nodes, initNode)
				} else {
					if tokens[i+2].Type_ == lexer.LITERAL {
						literalNode := createLiteralNode(tokens[i+2].Lexeme)
						initNode := createInitializationExpressionNode(id, literalNode)
						nodes = append(nodes, initNode)
					} else if tokens[i+2].Type_ == lexer.IDENTIFIER {
						literalNode := createLiteralNode(tokens[i+2].Lexeme)
						initNode := createInitializationExpressionNode(id, literalNode)
						nodes = append(nodes, initNode)
					} else {
						invalidInitValue()
					}
				}
			} else {
				invalidAssignment()
			}

			count++
		} else {
			if !inStatement {
				invalidStatement()
			}
		}
	}

	programNode := createProgramNode(nodes)
	fmt.Println(programNode)
}
