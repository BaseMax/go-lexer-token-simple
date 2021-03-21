package main

// struct Position
type Position struct {
	line   int
	column int
}

// Token data type is a int
type Token int

const (
	EOF = iota + 1 // start from 1
	UNKNOWN // so 2, etc...

	// keywords
	IDENT // [a-zA-Z0-9]+
	INT   // int
	RET   // ret (AKA return)

	// operators
	ADD // +
	SUB // -
	MUL // *
	DIV // /

	ASSIGN // =
	SEMI   // ;

	PARENT_LEFT // (
	PARENT_RIGHT // )

	SECTION_LEFT  // {
	SECTION_RIGHT // }
)

var tokens = []string{
	UNKNOWN: "UNKNOWN",
	EOF:     "EOF",

	// keywords
	IDENT:   "IDENT",
	INT:     "INT",
	RET:     "RET",

	// operators
	ADD: 	 "+",
	SUB: 	 "-",
	MUL: 	 "*",
	DIV: 	 "/",

	ASSIGN:  "=",
	SEMI:    ";",

	PARENT_LEFT:  "(",
	PARENT_RIGHT: ")",

	SECTION_LEFT: "{",
	SECTION_RIGHT: "}",
}

// Returns name of a token from it's number (Token data type)
func (t Token) String() string {
	return tokens[t]
}
