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

func PrecedenceOf(tt token.TokenType) int {
	if 0 <= tt && int(tt) < len(precedences) {
		return precedences[tt]
	}

	return PrecLowest
}
