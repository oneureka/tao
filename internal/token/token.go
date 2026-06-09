package token

import (
	"strconv"
	"strings"
)

type Token struct {
	Type   TokenType
	Lexeme string
	Value  any
	Span   Span
}

func (tok Token) literal() any {
	switch tok.Type {
	case Int:
		value, _ := strconv.Atoi(tok.Lexeme)
		return value
	case Float:
		value, _ := strconv.ParseFloat(tok.Lexeme, 64)
		return value
	case String:
		return tok.Lexeme[1 : len(tok.Lexeme)-2]
	case Bool:
		return strings.HasPrefix(tok.Lexeme, "T")
	default:
		return nil
	}
}

func New(_type TokenType, lexeme string, span Span) Token {
	token := Token{_type, lexeme, nil, span}
	token.Value = token.literal()
	return token
}
