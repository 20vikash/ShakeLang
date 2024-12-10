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

type BinaryExpression struct {
	_type    string
	left     Node
	operator string
	right    Node
}

type VariableDeclaration struct {
	_type       string
	declaration VariableDeclarator
}

type VariableDeclarator struct {
	_type string
	id    Identifier
	init  Literal
}

type Literal struct {
	_type string
	value string
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

func (p VariableDeclarator) iHaveTo() {}

func (p Literal) iHaveTo() {}
