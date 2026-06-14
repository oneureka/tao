package lexical

import "github.com/oneureka/tao/internal/token"

func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		cursor: NewCursor(source),
		tokens: make([]token.Token, 0),
		single: false,
	}
}

func NewCursor(input string) *Cursor {
	cursor := &Cursor{input: input}
	cursor.Line = 1
	cursor.Col = 1
	return cursor
}
