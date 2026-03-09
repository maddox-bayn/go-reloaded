package processor

func isPunctuation(c rune) bool {
	switch c {
	case '.', ',', ';', ':', '!', '?':
		return true
	}
	return false

}

func Tokenize(text string) []string {
	var result []string
	runes := []rune(text)
	word := ""
	insideBrackeks := false

	for i := 0; i < len(runes); i++ {
		char := runes[i]
		switch {
		case insideBrackeks:
			word += string(char)
			if char == ')' {

				insideBrackeks = false
			}
		case char == '(':
			if word != "" {
				result = append(result, word)
				word = ""
			}
			word += string(char)
			insideBrackeks = true
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
			continue
		case isPunctuation(char):
			if word != "" {
				result = append(result, word)
				word = ""
			}
			start := i
			for i+1 < len(runes) && isPunctuation(runes[i+1]) {
				i++
			}
			token := string(runes[start : i+1])
			result = append(result, token)
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
