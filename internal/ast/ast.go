package ast

import "github.com/oneureka/tao/internal/token"

type Node interface {
	Pos() int
	End() int
}

type Expr interface {
	Acceptable
	Node
}

type Statement interface {
	Acceptable
	Node
}

type Declaration interface {
	Acceptable
	Node
}

type IfStatement struct {
	Cond Expr
	Then BlockStatement
	Else BlockStatement
	Statement
}

type LoopStatement struct {
	Body BlockStatement
	Statement
}

type LeaveStatement struct {
	Keyword token.Token
	Statement
}

type ExprStatement struct {
	Expression Expr
	Statement
}

type FunctionDecl struct {
	Name     Identifier
	Receiver *Identifier
	Params   []token.Token
	Body     BlockStatement
	Declaration
}

type FunctionExpr struct {
	Params []token.Token
	Body   BlockStatement
	Expr
}

type BlockStatement struct {
	Statements []Statement
	Statement
}

type ReturnStatement struct {
	Keyword token.Token
	Value   Expr
	Statement
}

type VarDeclaration struct {
	Name        token.Token
	Mutable     bool
	Initializer Expr
	Declaration
}

type AssignExpr struct {
	Left     Expr
	Operator token.Token
	Right    Expr
	Expr
}

type Identifier struct {
	Name token.Token
	Expr
}

type BinaryExpr struct {
	Left     Expr
	Operator token.Token
	Right    Expr
	Expr
}

type UnaryExpr struct {
	Operator token.Token
	Right    Expr
	Expr
}

type CallExpr struct {
	Callee Expr
	Paren  token.Token
	Args   []Expr
	Expr
}

type Literal struct {
	Type  token.Token
	Value any
	Expr
}
