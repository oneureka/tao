package lexical

import (
	"unicode/utf8"

	"github.com/oneureka/tao/internal/token"
)

type Scanner struct {
	source   string
	cursor   *Cursor
	tokens   []token.Token
	start    int
	startPos token.Position
	nextTok  token.Token
	single   bool
	*ErrorSink
}

func (s *Scanner) NextToken() token.Token {
	s.single = true
	s.ScanTokens()

	s.single = false
	return s.nextTok
}

func (s *Scanner) ScanTokens() []token.Token {
	for {
		if s.cursor.EOF() {
			if s.nextTok.Type != token.EOF {
				s.nextTok = s.addToken(token.EOF, "")
			}

			break
		}

		s.start = s.cursor.Start()
		s.startPos = s.cursor.Position

		tok, ok := s.scanToken()

		if ok {
			s.nextTok = tok

			if s.single {
				return []token.Token{}
			}
		}
	}

	return s.tokens
}

func (s *Scanner) scanToken() (token.Token, bool) {
	c := s.cursor.Advance()

	if s.skipWhitespace(c) {
		return token.Token{}, false
	}

	if s.skipComment(c) {
		return token.Token{}, false
	}

	switch c {
	case '+':
		s.addToken(token.Plus, "+")
	case '*':
		s.addToken(token.Star, "*")
	case '/':
		s.addToken(token.Slash, "/")
	case '%':
		s.addToken(token.Modulo, "%")
	case '{':
		s.addToken(token.LBrace, "{")
	case '}':
		s.addToken(token.RBrace, "}")
	case '[':
		s.addToken(token.LSquare, "[")
	case ']':
		s.addToken(token.RSquare, "]")
	case '(':
		s.addToken(token.LParen, "(")
	case ')':
		s.addToken(token.RParen, ")")
	case ',':
		s.addToken(token.Comma, ",")
	case ':':
		s.addToken(token.Colon, ":")
	case ';':
		s.addToken(token.Semi, ";")
	case '.':
		s.addToken(token.Dot, ".")
	case '>':
		if s.match('=') {
			s.addToken(token.GreaterEq, ">=")
		} else {
			s.addToken(token.Greater, ">")
		}
	case '<':
		if s.match('=') {
			s.addToken(token.LessEq, "<=")
		} else {
			s.addToken(token.Less, "<")
		}
	case '=':
		if s.match('=') {
			s.addToken(token.EqualEq, "==")
		} else {
			if s.match('>') {
				s.addToken(token.EqArrow, "=>")
			} else {
				s.addToken(token.Equal, "=")
			}
		}
	case '!':
		if s.match('=') {
			s.addToken(token.BangEq, "!=")
		} else {
			s.addToken(token.Bang, "!")
		}
	case '-':
		if s.match('>') {
			s.addToken(token.Arrow, "->")
		} else {
			s.addToken(token.Minus, "-")
		}
	case '?':
		s.addToken(token.Illegal, "")
	case '&':
		if s.match('&') {
			if s.match('=') {
				s.addToken(token.Illegal, "&&=")
			} else {
				s.addToken(token.And, "&&")
			}
		} else {
			s.addToken(token.Illegal, "&")
		}
	case '|':
		if s.match('|') {
			if s.match('=') {
				s.addToken(token.Illegal, "||=")
			} else {
				s.addToken(token.Or, "||")
			}
		} else {
			if s.match('>') {
				s.addToken(token.Illegal, "|>")
			} else {
				s.addToken(token.Illegal, "|")
			}
		}
	case '"':
		if !s.stringToken() {
			s.addError(UnterminatedString, s.startPos)
			s.addToken(token.Illegal, "")
		}
	default:
		if !s.numberToken(c) {
			if !s.identifierToken(c) {
				s.addError(UnexpectedChar, s.startPos)
				s.addToken(token.Illegal, "")
			}
		}
	}

	return s.tokens[len(s.tokens)-1], true
}

func (s *Scanner) skipWhitespace(c rune) bool {
	if c == '\n' {
		s.cursor.Line += 1
		s.cursor.Col = 1

		return true
	}

	return isWhitespace(c)
}

func (s *Scanner) skipComment(c rune) bool {
	if c == '/' && s.match('/') {
		for {
			if s.cursor.Peek() == '\n' {
				break
			}

			if s.cursor.EOF() {
				break
			}

			s.cursor.Advance()
		}

		return true
	}

	return false
}

func (s *Scanner) identifierToken(c rune) bool {
	if isAlpha(c) {
		for isAlphaNumeric(s.cursor.Peek()) {
			s.cursor.Advance()
		}

		lexeme := s.cursor.Lexeme()

		switch {
		case isBool(lexeme):
			s.addToken(token.Bool, lexeme)
		case isNone(lexeme):
			s.addToken(token.Identifier, lexeme)
		default:
			tt, ok := token.Keywords[lexeme]

			if !ok {
				tt = token.Identifier
			}

			s.addToken(tt, lexeme)
		}

		return true
	}

	return false
}

func (s *Scanner) stringToken() bool {
	for {
		if s.cursor.Peek() == '"' {
			break
		}

		if s.cursor.EOF() {
			break
		}

		if s.cursor.Peek() == '\n' {
			s.cursor.Advance()
			s.cursor.Line += 1
			s.cursor.Col = 1
			break
		}

		s.cursor.Advance()
	}

	if s.cursor.Col == 1 {
		return false
	}

	if s.cursor.EOF() {
		return false
	}

	s.cursor.Advance()
	lexeme := s.cursor.Lexeme()

	s.addToken(token.String, lexeme)
	return true
}

func (s *Scanner) numberToken(c rune) bool {
	if isDigit(c) {
		tt := token.Int

		for isDigit(s.cursor.Peek()) {
			s.cursor.Advance()
		}

		if s.cursor.Peek() == '.' && isDigit(s.cursor.PeekNext()) {
			tt = token.Float
			s.cursor.Advance()

			for isDigit(s.cursor.Peek()) {
				s.cursor.Advance()
			}
		}

		lexeme := s.cursor.Lexeme()
		s.addToken(tt, lexeme)

		return true
	}

	return false
}

func (s *Scanner) addToken(tt token.TokenType, lexeme string) token.Token {
	tok := token.Token{
		Type:   tt,
		Lexeme: lexeme,
		Range: token.Range{
			Start: s.startPos,
			End:   s.cursor.Position,
		},
	}

	s.tokens = append(s.tokens, tok)
	return tok
}

func (s *Scanner) match(c rune) bool {
	if s.cursor.EOF() {
		return false
	}

	if s.cursor.Peek() == c {
		s.cursor.Forward(utf8.RuneLen(c))
		return true
	}

	return false
}
