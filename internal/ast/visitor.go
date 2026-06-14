package ast

type Acceptable interface {
	Accept(v Visitor) any
}

type Visitor interface{}

type ExprVisitor interface {
	Visitor
}
