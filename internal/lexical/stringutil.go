package lexical

const eof = rune(-1)

func isBool(str string) bool {
	return str == "True" || str == "False"
}

func isNone(str string) bool {
	return str == "None"
}

func isWhitespace(c rune) bool {
	return c == ' ' ||
		c == '\n' ||
		c == '\t' ||
		c == '\r'
}

func isAlphaNumeric(c rune) bool {
	return isAlpha(c) || isDigit(c)
}

func isAlpha(c rune) bool {
	return ('A' <= c && c <= 'Z') ||
		('a' <= c && c <= 'z') ||
		('_' == c)
}

func isDigit(c rune) bool {
	return '0' <= c && c <= '9'
}
