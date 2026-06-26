package processor

import "strings"

func isPunctuation(c rune) bool {
	switch c {
	case '.', ',', ';', ':', '!', '?':
		return true
	}
	return false

}

// take the text as string and tokenize it into slice of string
// tokenizatoin works seperate modifiers, punctuation (,.?!;:) and also handle single quote
func Tokenize(text string) []string {
	var result []string
	runes := []rune(text)
	var tempStr strings.Builder
	insideBracket := false

	for i := 0; i < len(runes); i++ {
		char := runes[i]
		switch {
		case insideBracket:
			tempStr.WriteRune(char)
			if char == ')' {

				insideBracket = false
				if tempStr.Len() != 0 {
					result = append(result, tempStr.String())
					tempStr.Reset()
				}
			}
			continue
		case char == '(':
			if tempStr.Len() != 0 {
				result = append(result, tempStr.String())
				tempStr.Reset()
			}
			tempStr.WriteRune(char)
			insideBracket = true
			continue
			// handle single as a string
		case char == '\'':
			if tempStr.Len() != 0 {
				result = append(result, tempStr.String())
				tempStr.Reset()
			}
			result = append(result, "'")
			continue
		case char == ' ':
			if tempStr.Len() != 0 {
				result = append(result, tempStr.String())
				tempStr.Reset()
			}
			continue
			// handle punctuation and group punctuation as single string
		case isPunctuation(char):
			if tempStr.Len() != 0 {
				result = append(result, tempStr.String())
				tempStr.Reset()
			}
			start := i
			for i+1 < len(runes) && isPunctuation(runes[i+1]) {
				i++
			}
			token := string(runes[start : i+1])
			result = append(result, token)
			continue
		default:
			tempStr.WriteRune(char)
		}

	}
	if tempStr.Len() != 0 {
		result = append(result, tempStr.String())
		tempStr.Reset()
	}
	return result
}
