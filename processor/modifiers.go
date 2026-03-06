package processor

import (
	"strconv"
	"strings"
)

// ["this","is","great","(up,2)"]

func ApplyModifiers(tokens []string) []string {
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		if strings.HasPrefix(token, "(") && strings.HasSuffix(token, ")") {

			content := token[1 : len(token)-1]

			part := strings.Split(content, ",")

			cmd := strings.TrimSpace(part[0])

			count := 1

			if len(part) > 1 {

				n, err := strconv.Atoi(strings.TrimSpace(part[1]))
				if err == nil {
					count = n
				}
			}
			switch cmd {
			case "hex":
				if i > 0 {
					num, err := strconv.ParseInt(tokens[i-1], 16, 64)
					if err == nil {
						tokens[i-1] = strconv.FormatInt(num, 10)
					}

				}
				//tokens = append(tokens[:i], tokens[i+1:]...)
			case "bin":
				if i > 0 {
					num, err := strconv.ParseInt(tokens[i-1], 2, 64)
					if err == nil {
						tokens[i-1] = strconv.FormatInt(num, 10)
					}
					// -tokens = append(tokens[:i], tokens[i+1:]...)
				}
			case "up":
				for j := 1; j <= count && i-j >= 0; j++ {
					tokens[i-j] = strings.ToUpper(tokens[i-j])
					//tokens = append(tokens[:i], tokens[i+1:]...)
				}
			case "low":
				for j := 1; j <= count && i-j >= 0; j++ {
					tokens[i-j] = strings.ToLower(tokens[i-j])
					//okens = append(tokens[:i], tokens[i-1:]...)
				}
			case "cap":
				for j := 1; j <= count && i-j >= 0; j++ {
					rstr := []rune(tokens[i-j])
					if len(rstr) > 0 && rstr[0] >= 'a' && rstr[0] <= 'z' {
						rstr[0] -= 32
					}
					tokens[i-j] = string(rstr)
					//tokens = append(tokens[:i], tokens[i+1:]...)
				}
			}
			tokens = append(tokens[:i], tokens[i+1:]...)
			i--
		}

	}
	return tokens
}
