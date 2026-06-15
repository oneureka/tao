package parsing

import "github.com/oneureka/tao/internal/token"

const (
	PrecLowest int = iota
	PrecAssign
	PrecPipe
	PrecOr
	PrecAnd
	PrecEquality
	PrecCompare
	PrecAdd
	PrecMul
	PrecUnary
	PrecCall
)

var precedences = [...]int{
	token.Plus:      PrecAdd,
	token.Minus:     PrecAdd,
	token.Star:      PrecMul,
	token.Slash:     PrecMul,
	token.Modulo:    PrecMul,
	token.Greater:   PrecCompare,
	token.GreaterEq: PrecCompare,
	token.Less:      PrecCompare,
	token.LessEq:    PrecCompare,
	token.Equal:     PrecAssign,
	token.EqualEq:   PrecEquality,
	token.BangEq:    PrecEquality,
	token.And:       PrecAnd,
	token.Or:        PrecOr,
}

func PrecedenceOf(tok token.Token) int {
	if 0 <= tok.Type && int(tok.Type) < len(precedences) {
		return precedences[tok.Type]
	}

	return PrecLowest
}
