package ast

import "github.com/oneureka/tao/internal/token"

type DataDeclaration struct {
	Name   token.Token
	Fields []token.Token
	Declaration
}

type Self struct {
	Expr
}
