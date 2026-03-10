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
	insideQuote := false

	var r strings.Builder
	prev := ""

	for _, token := range tokens {
		if token == "" {
			continue
		}

		if prev != "" {
			switch {
			case Punctuation(token):

			case isQuote(token) && !insideQuote:
				if !Punctuation(prev) {
					r.WriteRune(' ')
				}

			case insideQuote:

			case isQuote(token) && insideQuote:

			case prev == "'" && !insideQuote:
				r.WriteRune(' ')

			default:
				r.WriteRune(' ')
			}
		}
		r.WriteString(token)
		prev = token

		if isQuote(token) {
			insideQuote = !insideQuote
		}

	}
	return r.String()
}
