package parsing

import "github.com/oneureka/tao/internal/token"

type Rule struct {
	Prefix string
	Infix  string
}

var rules = [...]Rule{
	token.Plus:       {Infix: "parse_binary"},
	token.Minus:      {Prefix: "parse_unary", Infix: "parse_binary"},
	token.Star:       {Infix: "parse_binary"},
	token.Slash:      {Infix: "parse_binary"},
	token.Modulo:     {Infix: "parse_binary"},
	token.Greater:    {Infix: "parse_compare"},
	token.GreaterEq:  {Infix: "parse_compare"},
	token.Less:       {Infix: "parse_compare"},
	token.LessEq:     {Infix: "parse_compare"},
	token.Equal:      {Infix: "parse_assign"},
	token.EqualEq:    {Infix: "parse_infix"},
	token.BangEq:     {Infix: "parse_infix"},
	token.And:        {Infix: "parse_binary"},
	token.Or:         {Infix: "parse_binary"},
	token.Not:        {Prefix: "parse_unary"},
	token.LBrace:     {Prefix: "parse_prefix"},
	token.LSquare:    {Prefix: "parse_array", Infix: "parse_index"},
	token.LParen:     {Prefix: "parse_grouping", Infix: "parse_call"},
	token.Identifier: {Prefix: "parse_identifier"},
	token.String:     {Prefix: "parse_literal"},
	token.Float:      {Prefix: "parse_literal"},
	token.Int:        {Prefix: "parse_literal"},
	token.Bool:       {Prefix: "parse_literal"},
	token.Dot:        {Infix: "parse_infix"},
	token.If:         {Prefix: "parse_if"},
	token.Fun:        {Prefix: "parse_func"},
}
