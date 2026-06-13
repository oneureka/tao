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
	*ErrorSink
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
