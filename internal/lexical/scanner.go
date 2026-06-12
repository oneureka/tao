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
