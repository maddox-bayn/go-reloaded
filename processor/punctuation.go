package processor

import (
	"strings"
)

// helper function to check if string is a qoute
func isQoute(s string) bool {
	return s == "'"
}

// helper function to check the project specified punction
func Punctuation(c string) bool {
	switch rune(c[0]) {
	case '.', ',', ';', ':', '!', '?':
		return true
	}
	return false
}

// Rebuild the tokens into a string of transformed and modified text

func Rebuild(tokens []string) string {

	var b strings.Builder

	insideQoute := false

	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if i-1 >= 0 {

			isClosingQoute := token == "'" && insideQoute
			isprevOpenQoute := tokens[i-1] == "'" && insideQoute

			if !Punctuation(token) && !isClosingQoute && !isprevOpenQoute {
				b.WriteRune(' ')
			}

		}

		if isQoute(token) {
			insideQoute = !insideQoute
		}
		b.WriteString(tokens[i])
	}
	return b.String()
}
