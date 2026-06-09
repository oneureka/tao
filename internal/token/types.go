package token

type TokenType int

const (
	Illegal TokenType = iota
	Identifier
	String
	Float
	Int
	Bool

	Plus
	Minus
	Star
	Slash
	Modulo

	Greater
	GreaterEq
	Less
	LessEq
	Equal
	EqualEq
	Bang
	BangEq

	And
	Or
	Not
	Pipe

	LBrace
	RBrace
	LSquare
	RSquare
	LParen
	RParen

	Comma
	Colon
	Semi
	Dot

	If
	Else
	Loop
	Leave
	Switch
	Package
	Import
	Const
	Data
	Fun
	Let
	Mut
	Try
	Return

	EqArrow
	Arrow
	EOF
)
