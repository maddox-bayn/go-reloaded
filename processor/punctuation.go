package processor

import (
	"strings"
)

func Punctuation(c rune) bool {
	switch c {
	case '.', ',', ';', ':', '!', '?':
		return true
	}
	return false

}

func Rebuild(tokens []string) string {
	//insideQuote := false
	var r strings.Builder
	prev := ""

	for _, token := range tokens {

		//if i > 0 && !Punctuation(rune(token[0])) && token != "'" && tokens[i-1] != "'" {}
		if prev != "" &&
			!Punctuation(rune(token[0])) &&
			token != "'" &&
			prev != "'" {
			r.WriteRune(' ')
		}

		r.WriteString(token)
		prev = token

	}
	return r.String()
}
