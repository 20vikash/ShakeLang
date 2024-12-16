package backend

import (
	"fmt"
	"shake/shake/frontend/parser"
	"strconv"
)

func Exec(ast []parser.Node) {
	variables := make(map[string]string)

	for _, v := range ast {
		if v.GetType() == "VariableDeclaration" {
			node := v.(parser.VariableDeclaration)
			_, exists := variables[node.Id.Name]

			if exists {
				break //TODO: Variable already exists.
			}

			if node.Init == nil {
				variables[node.Id.Name] = "0"
			} else {
				if node.Init.GetType() == "Identifier" {
					d, exists := variables[node.Init.(parser.Identifier).Name]

					if !exists {
						break //TODO: No such variable.
					}

					variables[node.Id.Name] = d
				} else if node.Init.GetType() == "Literal" {
					variables[node.Id.Name] = node.Init.(parser.Literal).Value
				} else if node.Init.GetType() == "BinaryExpression" {
					binex := node.Init.(parser.BinaryExpression)
					result := EvaluateBinaryExpression(binex.Left, binex.Right, binex.Operator, variables)

					variables[node.Id.Name] = strconv.Itoa(result)
				}
			}
		} else if v.GetType() == "InitializationExpression" {
			node := v.(parser.InitializationExpression)
			_, exists := variables[node.Id.Name]

			if !exists {
				break //TODO: Didn't declare the variable.
			}

			if node.Init.GetType() == "Identifier" {
				d, exists := variables[node.Init.(parser.Identifier).Name]

				if !exists {
					break //TODO: No such variable.
				}

				variables[node.Id.Name] = d
			} else if node.Init.GetType() == "Literal" {
				variables[node.Id.Name] = node.Init.(parser.Literal).Value
			} else if node.Init.GetType() == "BinaryExpression" {
				binex := node.Init.(parser.BinaryExpression)
				result := EvaluateBinaryExpression(binex.Left, binex.Right, binex.Operator, variables)

				variables[node.Id.Name] = strconv.Itoa(result)
			}
		} else if v.GetType() == "ProclaimStatement" {
			node := v.(parser.ProclaimStatement)

			if node.Arg.GetType() == "Identifier" {
				d, exists := variables[node.Arg.(parser.Identifier).Name]

				if !exists {
					break //TODO: Variable not declared.
				}

				fmt.Println(d)
			} else if node.Arg.GetType() == "Literal" {
				fmt.Println(node.Arg.(parser.Literal).Value)
			} else if node.Arg.GetType() == "BinaryExpression" {
				binex := node.Arg.(parser.BinaryExpression)
				result := EvaluateBinaryExpression(binex.Left, binex.Right, binex.Operator, variables)

				fmt.Println(result)
			}
		}
	}
}
