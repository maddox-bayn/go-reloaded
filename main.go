package main

import (
	"fmt"
	//"strings"
	"go-reloaded/processor"
)

func main() {
	/*
		a := "hello  ,world! (up, 2) hello !!! (up, 2);;; hello(up) 'awesome' ... Simply add 42 (hex) and 10 (bin)"
		c := processor.Tokenize(a)
		fmt.Println(processor.ApplyModifiers(c))
	*/
	b := []string{"boy", "is", "man", "A", ""}
	fmt.Println(processor.FixArticles(b))
}
