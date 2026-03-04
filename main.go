package main

import (
	"fmt"
	"strconv"
	//"go-reloaded/processor"
)

func main() {
	//a := "hello, world! (up, 2) hello !!! (up, 2);;; hello(up) ' awesome ' ... Simply add 42 (hex) and 10 (bin)"
	//c := processor.Tokenize(a)
	//fmt.Println(c)
	tokens := []string{"42", "(hex)", "and", "FF", "(hex)"}

	for i, token := range tokens {
		switch token {
		case "(hex)":
			num, err := strconv.ParseInt(tokens[i-1], 16, 64) // ← one line
			if err != nil {
				continue // bad input? skip safely
			}
			tokens[i-1] = fmt.Sprintf("%d", num)
			tokens[i] = ""
		}
	}
	fmt.Println(tokens)
}
