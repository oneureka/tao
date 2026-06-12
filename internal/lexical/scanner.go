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

	if s.cursor.EOF() {
		s.addError(UnterminatedString, s.startPos)
		s.addToken(token.Illegal, "")
		return true
	}

	if s.cursor.Col == 1 {
		s.addError(UnterminatedString, s.startPos)
		s.addToken(token.Illegal, "")
		return true
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
