package processor

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// apply modificatoin to text using modifier provided in the text
func ApplyModifiers(tokens []string) []string {
	// a write index to mark transition
	writeIndex := 0
	for i := 0; i < len(tokens); i++ {
		token := tokens[i]
		// capturing modifier
		if strings.HasPrefix(token, "(") && strings.HasSuffix(token, ")") {

			content := token[1 : len(token)-1]

			part := strings.Split(content, ",")

			cmd := strings.TrimSpace(part[0])

			count := 1

		//  when count is provided in modifier
			if len(part) > 1 {
				n, err := strconv.Atoi(strings.TrimSpace(part[1]))
				if err == nil && n > count {
					count = n
				}
			}
			// mark where, the start of the begaining of the inner loop 
			start := writeIndex - count
			if start < 0 {
				start = 0
			}
			
			for j := start; j < writeIndex; j++ {
				switch strings.ToLower(cmd) {
				case "hex":
					num64, err := strconv.ParseInt(tokens[j], 16, 64)
					if err != nil {
						log.Fatal(err)
					}
					tokens[j] = strconv.FormatInt(num64, 10)
				case "bin":
					num64, err := strconv.ParseInt(tokens[j], 2, 64)
					if err != nil {
						log.Fatal(err)
					}
					tokens[j] = strconv.FormatInt(num64, 10)
				case "up":
					tokens[j] = strings.ToUpper(tokens[j])
				case "low":
					tokens[j] = strings.ToLower(tokens[j])
				case "cap":
					tokens[j] = strings.ToUpper(tokens[j][:1]) + strings.ToLower(tokens[j][1:])
				}
			}
			// skip each case when we hit and use modifier to override the index and value 
			continue
		}
		tokens[writeIndex] = tokens[i]
		writeIndex++
	}
	return tokens[:writeIndex]
}
