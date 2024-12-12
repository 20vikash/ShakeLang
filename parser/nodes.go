package parser

type Node interface {
	iHaveTo()
}

type Program struct {
	_type string
	body  []Node
}

type Identifier struct {
	_type string
	name  string
}

func createIdentifierNode(name string) Identifier {
	return Identifier{_type: "Identifier", name: name}
}

type BinaryExpression struct {
	_type    string
	left     Node
	operator string
	right    Node
}

func createBinaryExpressionNode(left Node, operator string, right Node) BinaryExpression {
	return BinaryExpression{_type: "BinaryExpression", left: left, right: right, operator: operator}
}

type VariableDeclaration struct {
	_type string
	id    Identifier
	init  Node
}

func createVariableDeclarationNode(id Identifier, init Node) VariableDeclaration {
	return VariableDeclaration{_type: "VariableDeclaration", id: id, init: init}
}

type Literal struct {
	_type string
	value string
}

func createLiteralNode(value string) Literal {
	return Literal{_type: "Literal", value: value}
}

type SubscriptExpression struct {
	_type     string
	object    Identifier
	subscript Node
}

func (p Program) iHaveTo() {}

func (p Identifier) iHaveTo() {}

func (p BinaryExpression) iHaveTo() {}

func (p VariableDeclaration) iHaveTo() {}

func (p Literal) iHaveTo() {}
