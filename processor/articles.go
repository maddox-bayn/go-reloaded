package processor

import "strings"

// fixarticles uses correct article to formate tokens
func FixArticles(tokens []string) []string {

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
