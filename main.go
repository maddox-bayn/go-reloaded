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
	tokens := []string{"ema", "(up)", "and", "ff", "(up)"}
	for i := 0; i < len(tokens); i++ {
		if tokens[i] == "(up)" && i > 0 {
			str := tokens[i-1]
			str = strings.ToUpper(str)
			tokens[i-1] = str
			tokens[i] = ""
		}
	}
	fmt.Println(tokens)
}
