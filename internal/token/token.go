package token

import (
	"fmt"
	"strconv"
	"strings"
)

type Token struct {
	Type   TokenType
	Lexeme string
	Value  any
	Range
}

type Range struct {
	Start Position
	End   Position
}

type Position struct {
	Offset int
	Line   int
	Col    int
}

func (tok Token) Literal() any {
	switch tok.Type {
	case Int:
		num, _ := strconv.Atoi(tok.Lexeme)
		return num
	case Float:
		num, _ := strconv.ParseFloat(tok.Lexeme, 64)
		return num
	case String:
		return tok.Lexeme[1 : len(tok.Lexeme)-2]
	case Bool:
		return strings.HasPrefix(tok.Lexeme, "T")
	default:
		return nil
	}
}

func (tok Token) String() string {
	return fmt.Sprintf("Token(%v, %s)", tok.Type, tok.Lexeme)
}
