package backend

import (
	"shake/shake/frontend/parser"
	"strconv"
)

func getSoloValue(element parser.Node, variables map[string]string) int {
	if element.GetType() == "Literal" {
		d, err := strconv.Atoi(element.(parser.Literal).Value)

		if err != nil {
			//TODO: Invalid type for arithmetic.
		}

		return d
	} else {
		d, exists := variables[element.(parser.Identifier).Name]

		if !exists {
			//TODO: Variable doesn't exist.
		}

		d1, err := strconv.Atoi(d)

		if err != nil {
			//TODO: Invalid type for arithmetic.
		}

		return d1
	}
}

func arithmetic(leftValue int, rightValue int, operator string) int {
	if operator == "+" {
		return leftValue + rightValue
	} else if operator == "-" {
		return leftValue - rightValue
	} else if operator == "*" {
		return leftValue * rightValue
	} else if operator == "/" {
		return leftValue / rightValue
	} else if operator == "%" {
		return leftValue % rightValue
	} else if operator == "^" {
		return leftValue ^ rightValue
	}

	return -1
}

func EvaluateBinaryExpression(left parser.Node, right parser.Node, operator string, variables map[string]string) int {
	if left.GetType() != "BinaryExpression" && right.GetType() != "BinaryExpression" {
		leftValue := getSoloValue(left, variables)
		rightValue := getSoloValue(right, variables)

		return arithmetic(leftValue, rightValue, operator)
	} else if (left.GetType() == "BinaryExpression") && (right.GetType() == "Literal" || right.GetType() == "Identifier") {
		rightValue := getSoloValue(right, variables)
		leftT := left.(parser.BinaryExpression)
		leftValue := EvaluateBinaryExpression(leftT.Left, leftT.Right, leftT.Operator, variables)

		return arithmetic(leftValue, rightValue, operator)
	} else if (right.GetType() == "BinaryExpression") && (left.GetType() == "Literal" || left.GetType() == "Identifier") {
		leftValue := getSoloValue(left, variables)
		rightT := right.(parser.BinaryExpression)
		rightValue := EvaluateBinaryExpression(rightT.Left, rightT.Right, rightT.Operator, variables)

		return arithmetic(leftValue, rightValue, operator)
	} else {
		leftT := left.(parser.BinaryExpression)
		rightT := right.(parser.BinaryExpression)
		leftValue := EvaluateBinaryExpression(leftT.Left, leftT.Right, leftT.Operator, variables)
		rightValue := EvaluateBinaryExpression(rightT.Left, rightT.Right, rightT.Operator, variables)

		return arithmetic(leftValue, rightValue, operator)
	}
}
