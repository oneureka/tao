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
	return p.parseLiteral()
}

func (p *Parser) parseLiteral() ast.Expr {
	tok := p.peek()
	tok.Value = tok.Literal()

	expr := ast.Literal{Type: tok, Value: tok.Value}
	p.advance()

	return expr
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
