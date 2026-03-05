package main

import (
	"fmt"
	"strings"
	//"go-reloaded/processor"
)

func main() {
	//a := "hello, world! (up, 2) hello !!! (up, 2);;; hello(up) ' awesome ' ... Simply add 42 (hex) and 10 (bin)"
	//c := processor.Tokenize(a)
	//fmt.Println(c)
	tokens := []string{"come", "ema", "and", "ff", "(up, 2)"}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(up, 2)" && i > 0 {
			for t := i - 1; t >= 0; t-- {
				str := tokens[t]
				str = strings.ToUpper(str)
				tokens[t] = str
				if t < 2 {
					break
				}
			}
			tokens = append(tokens[:i], tokens[i+1:]...)
		}
	}
	fmt.Println(tokens)
}
