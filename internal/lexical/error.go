package lexical

import "github.com/oneureka/tao/internal/token"

type LexicalError int

const (
	UnterminatedString LexicalError = iota
	UnexpectedChar
	IllegalToken
)

type ErrorSink struct {
	errors []Error
}

type Error struct {
	Code LexicalError
	Pos  token.Position
}

func (e *ErrorSink) addError(err LexicalError, pos token.Position) {
	e.errors = append(e.errors, Error{Code: err, Pos: pos})
}

func (e *ErrorSink) Errors() []Error {
	return e.errors
}
