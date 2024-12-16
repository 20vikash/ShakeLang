package parser

type Node interface {
	GetType() string
}

type Program struct {
	_type string
	Body  []Node
}

func createProgramNode(body []Node) Program {
	return Program{_type: "Program", Body: body}
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
	Id    Identifier
	Init  Node
}

func createVariableDeclarationNode(id Identifier, init Node) VariableDeclaration {
	return VariableDeclaration{_type: "VariableDeclaration", Id: id, Init: init}
}

func createVariableDeclarationNodeWithoutInit(id Identifier) VariableDeclaration {
	return VariableDeclaration{_type: "VariableDeclaration", Id: id}
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
	Id    Identifier
	Init  Node
}

type ProclaimStatement struct {
	_type string
	Arg   Node
}

func createProclaimStatementNode(arg Node) ProclaimStatement {
	return ProclaimStatement{_type: "ProclaimStatement", Arg: arg}
}

func createInitializationExpressionNode(id Identifier, init Node) InitializationExpression {
	return InitializationExpression{_type: "InitializationExpression", Id: id, Init: init}
}

func (p Identifier) GetType() string { return p._type }

func (p BinaryExpression) GetType() string { return p._type }

func (p InitializationExpression) GetType() string { return p._type }

func (p ProclaimStatement) GetType() string { return p._type }

func (p VariableDeclaration) GetType() string { return p._type }

func (p Literal) GetType() string { return p._type }
