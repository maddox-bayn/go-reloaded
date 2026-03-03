package internal

func Tokenize(text string) []string {
	var result []string
	runes := []rune(text)
	sep := rune(' ')
	coma := rune(',')
	dot := rune('.')
	quem := rune('!')
	opp := rune('(')
	word := ""
	IsInsideBrackecks := false

	for i := 0; i < len(runes)-1; i++ {
		char := runes[i]
		if char == opp {
			IsInsideBrackecks = true
		}
		switch {
		case IsInsideBrackecks:
			word += string(char)
		case char == ')':
			result = append(result, word)
			word = ""
			IsInsideBrackecks = false
		default:
			word += string(char)
		}
	}
}
