package lexical

import (
	"unicode/utf8"

	"github.com/oneureka/tao/internal/token"
)

type Cursor struct {
	input   string
	start   int
	current int
	token.Position
}

func (c *Cursor) Peek() rune {
	if c.EOF() {
		return eof
	}

	r, _ := utf8.DecodeRuneInString(c.input[c.current:])
	return r
}

func (c *Cursor) PeekNext() rune {
	if c.current+1 >= len(c.input) {
		return eof
	}

	r, _ := utf8.DecodeRuneInString(c.input[c.current+1:])
	return r
}

func (c *Cursor) Advance() rune {
	r, size := utf8.DecodeRuneInString(c.input[c.current:])
	c.current += size
	c.Offset += size
	c.Col += size
	return r
}

func (c *Cursor) Lexeme() string {
	return c.input[c.start:c.current]
}

func (c *Cursor) Start() int {
	c.start = c.current
	return c.start
}

func (c *Cursor) EOF() bool {
	return c.current >= len(c.input)
}
