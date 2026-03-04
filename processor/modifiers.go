package processor

import (
	"strconv"
)
/*
for i := 0; i < len(tokens); i++ {
    if tokens[i] == "(hex)" {
        // modify tokens[i-1]
        // remove tokens[i]
        tokens = append(tokens[:i], tokens[i+1:]...)
        i-- // IMPORTANT
    }
}
	*/

func ApplyModifiers(token []string) []string {
	for i := 0; i < len(token); i++ {
		switch {
		case token[i] == ("hex"):
			num, err := strconv.ParseInt(token[i-1], 16, 64)
			if err != nil {
				continue
			}
			token[i-1] = 
		}
	}
}
