package processor

import "strings"

func FixArticles(tokens []string) []string {

	//vowels := map[rune]bool{'a': true, 'e': true, 'i': true, 'u': true, 'o': true, 'h': true, 'A': true, 'E': true, 'I': true, 'O': true, 'U': true, 'H': true}

	for i := 0; i < len(tokens); i++ {

		if (tokens[i] == "a" || tokens[i] == "A") && i < len(tokens) {

			nextWord := tokens[i+1]

			if len(nextWord) == 0 {
				continue
			}
			if strings.ContainsRune("aeiouhAEIOUH", rune(nextWord[0])) {

				if tokens[i] == "a" {
					tokens[i] = "an"
				} else {
					tokens[i] = "An"
				}
			}
		}
	}
	return tokens
}
