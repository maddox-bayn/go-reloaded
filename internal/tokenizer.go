package internal

func isPunctuation(c rune) bool {
	switch c {
	case '.', ',', ';', ':':
		return true
	}
	return false
}

func Tokenize(text string) []string {
	var result []string
	runes := []rune(text)
	word := ""
	IsInsideBrackecks := false
	//IsInsideQuote := false

	for i := 0; i < len(runes); i++ {
		char := runes[i]
		switch {
		case IsInsideBrackecks:
			word += string(char)
			if char == ')' {
				result = append(result, word)
				word = ""
				IsInsideBrackecks = false
				continue
			}
		case char == '(':
			if word != "" {
				result = append(result, word)
				word = ""
			}
			word += string(char)
			IsInsideBrackecks = true
			continue
		case char == '\'':
			if word != "" {
				result = append(result, word)
				word = ""
			}
			result = append(result, "'")
			continue
		case char == ' ':
			if word != "" {
				result = append(result, word)
				word = ""
			}
		case isPunctuation(char):
			if word != "" {
				result = append(result, word)
				word = ""
			}
			result = append(result, string(char))
			continue
		default:
			word += string(char)
		}

	}

	if word != "" {
		result = append(result, word)
		//word = ""
	}
	return result
}
