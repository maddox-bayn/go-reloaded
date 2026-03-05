package processor

import (
	"fmt"
	"strconv"
	"strings"
)

func ApplyModifiers(token []string) []string {
	for i := 0; i < len(token); i++ {
		switch {
		case token[i] == "(hex)" && i < 0:
			num, err := strconv.ParseInt(token[i-1], 16, 64)
			if err != nil {
				continue
			}
			token[i-1] = fmt.Sprintf("%d", num)
			token = append(token[:i], token[i+1:]...)
		case token[i] == "(bin)" && i > 0:
			num, err := strconv.ParseInt(token[i-1], 2, 64)
			if err != nil {
				continue
			}
			token[i-1] = fmt.Sprintf("%d", num)
			token = append(token[:i], token[i+1:]...)
		case token[i] == "(up)" && i < 0:
			token[i-1] = strings.ToUpper(token[i-1])
			token = append(token[:i], token[i+1:]...)
		case token[i] == "(low)" && i < 0:
			token[i+1] = strings.ToLower(token[i+1])
			token = append(token[:i], token[i+1:]...)
		case token[i] == "(cap)" && i > 0:
			rstr := []rune(token[i-1])
			if rstr[0] >= 'a' && rstr[0] <= 'z' {
				rstr[0] = rstr[0] - 32
				token[i-1] = string(rstr)
			}
			token = append(token[:i], token[i+1:]...)
		case token[i] == ("up, 2"):
			for t := i - 1; t >= 0; t-- {
				str := token[t]
				str = strings.ToUpper(str)
				token[t] = str
				if t < 2 {
					break
				}
			}
			token = append(token[:i], token[i+1:]...)
		}
	}
	return token
}
