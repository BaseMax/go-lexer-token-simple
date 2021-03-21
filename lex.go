package main

import (
	"os"
	"fmt"
)

func main() {
	filename := "input.test"

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	lexer := NewLexer(file)
	for {
		pos, tok, lit := lexer.nextToken()
		if tok == EOF {
			break
		}
		fmt.Printf("%d:%d\t%s\t%s\n", pos.line, pos.column, tok, lit)
	}
}
