package parser

type Node interface {
	iHaveTo()
}

type Program struct {
	_type string
	body  []Node
}

func createProgramNode(body []Node) Program {
	return Program{_type: "Program", body: body}
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

func createVariableDeclarationNodeWithoutInit(id Identifier) VariableDeclaration {
	return VariableDeclaration{_type: "VariableDeclaration", id: id}
}

type Literal struct {
	_type string
	value string
}

func createLiteralNode(value string) Literal {
	return Literal{_type: "Literal", value: value}
}

type InitializationExpression struct {
	_type string
	id    Identifier
	init  Node
}

type proclaimStatement struct {
	_type string
	arg   Node
}

func createProclaimStatementNode(arg Node) proclaimStatement {
	return proclaimStatement{_type: "ProclaimStatement", arg: arg}
}

func createInitializationExpressionNode(id Identifier, init Node) InitializationExpression {
	return InitializationExpression{_type: "InitializationExpression", id: id, init: init}
}

func (p Program) iHaveTo() {}

func (p Identifier) iHaveTo() {}

func (p BinaryExpression) iHaveTo() {}

func (p InitializationExpression) iHaveTo() {}

func (p proclaimStatement) iHaveTo() {}

func (p VariableDeclaration) iHaveTo() {}

func (p Literal) iHaveTo() {}
