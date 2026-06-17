package parsing

import (
	"github.com/oneureka/tao/internal/ast"
	"github.com/oneureka/tao/internal/token"
)

type Rule struct {
	prefix string
	infix  string
}

func (r *Rule) ParsePrefix(p *Parser) ast.Expr {
	var expr ast.Expr

	switch r.prefix {
	case "parse_unary":
		expr = p.parseUnary()
	case "parse_array":
	case "parse_grouping":
	case "parse_identifier":
	case "parse_literal":
	case "parse_fun":

	}

	return expr
}

func (r *Rule) ParseInfix(p *Parser, left ast.Expr) ast.Expr {
	var expr ast.Expr

	switch r.infix {
	case
		"parse_binary",
		"parse_compare",
		"parse_equality":
		expr = p.parseBinary(left)
	case "parse_assign":
	case "parse_index":
	case "parse_call":

	}

	return expr
}

var rules = [...]Rule{
	token.Plus:       {infix: "parse_binary"},
	token.Minus:      {prefix: "parse_unary", infix: "parse_binary"},
	token.Star:       {infix: "parse_binary"},
	token.Slash:      {infix: "parse_binary"},
	token.Modulo:     {infix: "parse_binary"},
	token.Greater:    {infix: "parse_compare"},
	token.GreaterEq:  {infix: "parse_compare"},
	token.Less:       {infix: "parse_compare"},
	token.LessEq:     {infix: "parse_compare"},
	token.Equal:      {infix: "parse_assign"},
	token.EqualEq:    {infix: "parse_equality"},
	token.BangEq:     {infix: "parse_equality"},
	token.And:        {infix: "parse_binary"},
	token.Or:         {infix: "parse_binary"},
	token.Not:        {prefix: "parse_unary"},
	token.LBrace:     {prefix: "parse_prefix"},
	token.LSquare:    {prefix: "parse_array", infix: "parse_index"},
	token.LParen:     {prefix: "parse_grouping", infix: "parse_call"},
	token.Identifier: {prefix: "parse_identifier"},
	token.String:     {prefix: "parse_literal"},
	token.Float:      {prefix: "parse_literal"},
	token.Int:        {prefix: "parse_literal"},
	token.Bool:       {prefix: "parse_literal"},
	token.Dot:        {infix: "parse_infix"},
	token.If:         {prefix: "parse_if"},
	token.Fun:        {prefix: "parse_fun"},
}

func RuleOf(tok token.Token) Rule {
	if 0 <= tok.Type && int(tok.Type) < len(rules) {
		return rules[tok.Type]
	}

	return Rule{}
}
