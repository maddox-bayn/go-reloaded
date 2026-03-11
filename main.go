package main

import (
	"fmt"
	"go-reloaded/processor"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . sample.txt, result.txt")
		return
	}

	inputfile := os.Args[1]
	outputfile := os.Args[2]

	data, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}

	text := string(data)

	token := processor.Tokenize(text)
	token = processor.ApplyModifiers(token)
	token = processor.FixArticles(token)
	result := processor.Rebuild(token)

	if !strings.HasSuffix(result, "\n") {
		result += "\n"
	}

	err = os.WriteFile(outputfile, []byte(result+"\n"), 0644)

	if err != nil {
		fmt.Println("Error writing file", err)
		return
	}
}
