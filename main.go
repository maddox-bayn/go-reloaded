package main

import (
	"fmt"
	"go-reloaded/processor"
)

func main() {
	a := "hello, world! (up, 2) hello !!! (up, 2);;; hello(up) ' awesome ' ..."
	c := processor.Tokenize(a)
	fmt.Println(c)
}
