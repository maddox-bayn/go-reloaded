package processor

func Punctuation(c string) bool {
	switch c {
	case ".", ",", "?", ":", ";":
		return true
	}
	return false
}

func Rebuild(tokens []string) string {
	result := tokens[0]
	if len(tokens) == 0 {
		return ""
	}
	for i := 1; i < len(tokens); i++ {
		//result += tokens[i]
		if Punctuation(tokens[i]) {
			result += tokens[i]
		} else {
			result += " " + tokens[i]
		}
	}
	return result
}
