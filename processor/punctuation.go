package processor

import (
	"strings"
)

func isQuote(s string) bool {
	return s == "'"
}

func Punctuation(c string) bool {
	switch rune(c[0]) {
	case '.', ',', ';', ':', '!', '?':
		return true
	}
	return false
}

func Rebuild(tokens []string) string {
	var b strings.Builder

	for i, token := range tokens {

		if i > 0 &&
			!Punctuation(token) &&
			token != "'" &&
			tokens[i-1] != "'" {

			b.WriteRune(' ')
		}
		b.WriteString(token)
	}
	return b.String()
}
