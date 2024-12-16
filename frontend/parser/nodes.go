package parser

type Node interface {
	GetType() string
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
	Name  string
}

func createIdentifierNode(name string) Identifier {
	return Identifier{_type: "Identifier", Name: name}
}

type BinaryExpression struct {
	_type    string
	Left     Node
	Operator string
	Right    Node
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
	Value string
}

func createLiteralNode(value string) Literal {
	return Literal{_type: "Literal", Value: value}
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

func (p Identifier) GetType() string { return p._type }

func (p BinaryExpression) GetType() string { return p._type }

func (p InitializationExpression) GetType() string { return p._type }

func (p proclaimStatement) GetType() string { return p._type }

func (p VariableDeclaration) GetType() string { return p._type }

func (p Literal) GetType() string { return p._type }
