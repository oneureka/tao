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

func (p *Parser) parseProgram() ast.Expr {
	defer func() {
		if r := recover(); r != nil {
			if _, ok := r.(ParseError); ok {
				p.synchronize()
			}
		}
	}()

	return p.parseExpression(PrecLowest)
}

func (p *Parser) parseStatement() ast.Statement {
	switch {
	case p.match(token.If):
		return p.parseIfStatement()
	case p.match(token.Loop):
		return p.parseLoopStatement()
	case p.match(token.Leave):
		return p.parseExprStatement()
	case p.match(token.Return):
		return p.parseReturnStatement()
	default:
		return p.parseExprStatement()
	}
}

func (p *Parser) parseDeclaration() ast.Declaration {
	switch {
	case p.match(token.Data):
		return nil
	case p.match(token.Fun):
		return p.parseFunctionDecl()
	case p.match(token.Let):
		return p.parseVarDeclaration()
	default:
		return p.parseStatement()
	}
}

func (p *Parser) parseIfStatement() ast.Statement {
	cond := p.parseExpression(PrecLowest)
	then := p.parseBlockStatement().(ast.BlockStatement)

	var elseBranch ast.BlockStatement

	if p.match(token.Else) {
		elseBranch = p.parseBlockStatement().(ast.BlockStatement)
	}

	return ast.IfStatement{Cond: cond, Then: then, Else: elseBranch}
}

func (p *Parser) parseLoopStatement() ast.Statement {
	body := p.parseBlockStatement().(ast.BlockStatement)
	return ast.LoopStatement{Body: body}
}

func (p *Parser) parseBlockStatement() ast.Statement {
	statements := make([]ast.Statement, 0)

	for {
		if p.check(token.RBrace) {
			break
		}

		if p.eof() {
			break
		}

		statements = append(statements, p.parseStatement())
	}

	p.consume(token.RBrace, "")
	return ast.BlockStatement{Statements: statements}
}

func (p *Parser) parseReturnStatement() ast.Statement {
	var value ast.Expr
	keyword := p.previous()

	if !p.check(token.Semi) {
		value = p.parseExpression(PrecLowest)
	}

	p.consume(token.Semi, "")
	return ast.ReturnStatement{Keyword: keyword, Value: value}
}

func (p *Parser) parseFunctionDecl() ast.Declaration {
	name := p.consume(token.Identifier, "")
	p.consume(token.LParen, "")

	params := make([]token.Token, 0)

	if !p.check(token.RParen) {
		for {
			if len(params) <= 128 {
				params = append(params, p.consume(token.Identifier, ""))

				if !p.match(token.Comma) {
					break
				}
			}

			panic(ParseError{token: p.peek()})
		}
	}

	p.consume(token.RParen, "")
	p.consume(token.LBrace, "")

	body := p.parseBlockStatement().(ast.BlockStatement)

	return ast.FunctionDecl{
		Name:   ast.Identifier{Name: name},
		Params: params,
		Body:   body,
	}
}

func (p *Parser) parseVarDeclaration() ast.Declaration {
	mut := p.match(token.Mut)
	tok := p.consume(token.Identifier, "")

	var initializer ast.Expr

	if p.match(token.Equal) {
		initializer = p.parseExpression(PrecLowest)

		return ast.VarDeclaration{Name: tok, Mutable: mut, Initializer: initializer}
	}

	panic(ParseError{})
}

func (p *Parser) parseExprStatement() ast.Statement {
	expr := p.parseExpression(PrecLowest)
	p.consume(token.Semi, "")

	return ast.ExprStatement{Expression: expr}
}

func (p *Parser) parseExpression(prec int) ast.Expr {
	rule := RuleOf(p.peek().Type)

	if rule.prefix == "" {
		panic(ParseError{})
	}

	left := rule.ParsePrefix(p)

	for prec < PrecedenceOf(p.peek().Type) {
		rule := RuleOf(p.peek().Type)

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

func (p *Parser) parseAssign(left ast.Expr) ast.Expr {
	operator := p.peek()
	p.advance()

	right := p.parseExpression(PrecAssign - 1)
	var expr ast.Expr

	switch lhs := left.(type) {
	case ast.Identifier:
		expr = ast.AssignExpr{Left: lhs, Operator: operator, Right: right}
	default:
		panic(ParseError{token: operator})
	}

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

	right := p.parseExpression(PrecedenceOf(p.previous().Type))
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

func (p *Parser) synchronize() {
	p.advance()

	for {
		if p.eof() {
			break
		}

		if p.previous().Type == token.Semi {
			return
		}

		switch p.peek().Type {
		case
			token.If,
			token.Loop,
			token.Switch,
			token.Data,
			token.Fun,
			token.Let,
			token.Return:
			return
		default:
			p.advance()
		}
	}
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
