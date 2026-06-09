package token

type Position struct {
	Filename string
	Offset   int
	Line     int
	Col      int
}

type Span struct {
	Start Position
	End   Position
}
