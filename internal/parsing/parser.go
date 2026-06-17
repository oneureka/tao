package parsing

import (
	"github.com/oneureka/tao/internal/ast"
	"github.com/oneureka/tao/internal/lexical"
	"github.com/oneureka/tao/internal/token"
)

type Parser struct {
	scanner *lexical.Scanner
	tokens  []token.Token
	current int
}

func (p *Parser) NextToken() token.Token {
	tok := p.scanner.NextToken()
	p.tokens = append(p.tokens, tok)

	return tok
}

func (p *Parser) parseExpression(prec int) ast.Expr {
	rule := RuleOf(p.peek())

	if rule.prefix == "" {
		panic(ParseError{})
	}

	left := rule.ParsePrefix(p)

	for prec < PrecedenceOf(p.peek()) {
		rule := RuleOf(p.peek())

		if rule.infix == "" {
			break
		}

		left = rule.ParseInfix(p, left)
	}

	return left
}

func (p *Parser) parseGrouping() ast.Expr {
	p.advance()
	expr := p.parseExpression(PrecLowest)

	p.expect(token.RParen)
	return expr
}

func (p *Parser) parseIdentifier() ast.Expr {
	expr := ast.Identifier{Name: p.peek()}
	p.advance()

	return expr
}

func (p *Parser) parseBinary(left ast.Expr) ast.Expr {
	operator := p.peek()
	p.advance()

	right := p.parseExpression(PrecedenceOf(p.previous()))
	return ast.BinaryExpr{Left: left, Operator: operator, Right: right}
}

func (p *Parser) parseUnary() ast.Expr {
	operator := p.peek()
	p.advance()

	right := p.parseExpression(PrecUnary)
	return ast.UnaryExpr{Operator: operator, Right: right}
}

func (p *Parser) parseLiteral() ast.Expr {
	tok := p.peek()
	tok.Value = tok.Literal()

	expr := ast.Literal{Type: tok, Value: tok.Value}
	p.advance()

	return expr
}

func (p *Parser) expect(tt token.TokenType) token.Token {
	return p.consume(tt, "")
}

func (p *Parser) consume(tt token.TokenType, message string) token.Token {
	if p.check(tt) {
		return p.advance()
	}

	panic(ParseError{token: p.peek(), message: message})
}

func (p *Parser) match(types ...token.TokenType) bool {
	for _, tt := range types {
		if p.check(tt) {
			p.advance()
			return true
		}
	}

	return false
}

func (p *Parser) check(tt token.TokenType) bool {
	if !p.eof() {
		return p.peek().Type == tt
	}

	return false
}

func (p *Parser) advance() token.Token {
	if !p.eof() {
		p.NextToken()
		p.current += 1
	}

	return p.previous()
}

func (p *Parser) previous() token.Token {
	return p.tokens[p.current-1]
}

func (p *Parser) peek() token.Token {
	return p.tokens[p.current]
}

func (p *Parser) eof() bool {
	return p.peek().Type == token.EOF
}
