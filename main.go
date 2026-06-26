package main

import (
	"fmt"
	"go-reloaded/processor"
	"os"
	"strings"
)

// engine and caller of the functions
func main() {
	// validat args and get file names
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . sample.txt, result.txt")
		return
	}
	inputfile := os.Args[1]
	outputfile := os.Args[2]
	// read the input file content and return it as byte
	data, err := os.ReadFile(inputfile)
	if err != nil {
		fmt.Println("Error reading file", err)
		return
	}
	// process file content through the processor
	text := string(data)

	token := processor.Tokenize(text)
	token = processor.ApplyModifiers(token)
	token = processor.FixArticles(token)
	result := processor.Rebuild(token)

	if !strings.HasSuffix(result, "\n") {
		result += "\n"
	}
	// write processed text to the output file
	err = os.WriteFile(outputfile, []byte(result), 0644)

	if err != nil {
		fmt.Println("Error writing file", err)
		return
	}
}
