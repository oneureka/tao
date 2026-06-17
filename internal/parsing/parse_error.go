package parsing

import (
	"fmt"

	"github.com/oneureka/tao/internal/token"
)

type ParseError struct {
	token   token.Token
	message string
}

func (p *ParseError) Error() string {
	if p.token.Type == token.EOF {
		return fmt.Sprintf("error at end, %s", p.message)
	}

	return fmt.Sprintf("error at line %d, %s", p.token.End.Line, p.message)
}
