package token

type TokenType int

const (
	Illegal TokenType = iota
	Identifier
	String
	Float
	Int
	Bool

	Plus
	Minus
	Star
	Slash
	Modulo

	Greater
	GreaterEq
	Less
	LessEq
	Equal
	EqualEq
	Bang
	BangEq

	And
	Or
	Not
	Pipe

	LBrace
	RBrace
	LSquare
	RSquare
	LParen
	RParen

	Comma
	Colon
	Semi
	Dot

	If
	Else
	Loop
	Leave
	Switch
	Package
	Import
	Const
	Data
	Fun
	Let
	Mut
	Try
	Return

	EqArrow
	Arrow
	EOF
)

func (tt TokenType) String() string {
	if 0 <= tt && int(tt) < len(tokens) {
		return tokens[tt]
	}

	return ""
}

var tokens = [...]string{
	Illegal:    "Illegal",
	Identifier: "Identifier",
	String:     "String",
	Float:      "Float",
	Int:        "Int",
	Bool:       "Bool",
	Plus:       "Plus",
	Minus:      "Minus",
	Star:       "Star",
	Slash:      "Slash",
	Modulo:     "Modulo",
	Greater:    "Greater",
	GreaterEq:  "GreaterEq",
	Less:       "Less",
	LessEq:     "LessEq",
	Equal:      "Equal",
	EqualEq:    "EqualEq",
	Bang:       "Bang",
	BangEq:     "BangEq",
	And:        "And",
	Or:         "Or",
	Not:        "Not",
	Pipe:       "Pipe",
	LBrace:     "LBrace",
	RBrace:     "RBrace",
	LSquare:    "LSquare",
	RSquare:    "RSquare",
	LParen:     "LParen",
	RParen:     "RParen",
	Comma:      "Comma",
	Colon:      "Colon",
	Semi:       "Semi",
	Dot:        "Dot",
	If:         "If",
	Else:       "Else",
	Loop:       "Loop",
	Leave:      "Leave",
	Switch:     "Switch",
	Package:    "Package",
	Import:     "Import",
	Const:      "Const",
	Data:       "Data",
	Fun:        "Fun",
	Let:        "Let",
	Mut:        "Mut",
	Try:        "Try",
	Return:     "Return",
	EqArrow:    "EqArrow",
	Arrow:      "Arrow",
	EOF:        "EOF",
}

var Keywords = map[string]TokenType{
	"if":      If,
	"else":    Else,
	"loop":    Loop,
	"leave":   Leave,
	"switch":  Switch,
	"package": Package,
	"import":  Import,
	"const":   Const,
	"data":    Data,
	"fun":     Fun,
	"let":     Let,
	"mut":     Mut,
	"try":     Try,
	"return":  Return,
	"and":     And,
	"or":      Or,
	"not":     Not,
}
