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
	insideQuote := false
	var r strings.Builder

	for i, token := range tokens {

		if token == "'" {
			if insideQuote == false {
				r.WriteString(" ")
				r.WriteString("'")
				insideQuote = true
				continue
			}
			r.WriteString("'")
			insideQuote = false
			continue
		}

		if i > 0 && !Punctuation(rune(token[0])) && token != "'" && tokens[i-1] != "'" {
			r.WriteString(" ")

		}
		r.WriteString(token)

	}
	return r.String()
}
