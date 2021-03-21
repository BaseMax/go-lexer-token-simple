package main

import (
	"io"
	"bufio"
	"unicode"
)

// struct of Lexer
type Lexer struct {
	position    Position
	reader 		*bufio.Reader
}

// Create a object/variable of Lexer and return address/pointer...
func NewLexer(reader io.Reader) *Lexer {
	return &Lexer{
		reader: bufio.NewReader(reader),
		position: Position{line: 1, column: 0},
	}
}

// nextToken scans the input for the next token.
func (l *Lexer) nextToken() (Position, Token, string) {
	// keep looping until we return a token
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return l.position, EOF, ""
			} else {
				// We don't know what should we do at this point!
				panic(err)
			}
		}

		// update the column to the position of the newly read in rune
		l.position.column++

		switch r {
		// new line
		// TODO: support \r\n
		case '\n':
			l.nextLine()

		// comments
		case '#':
			for {
				r, _, err := l.reader.ReadRune()
				if err != nil {
					if err == io.EOF {
						break
					}
				}

				if r == '\n' {
					l.nextLine()
					break
				} else {
					continue
				}
			}
			continue
		break

		// operators
		case '+':
			return l.position, ADD, ""
		case '-':
			return l.position, SUB, ""
		case '*':
			return l.position, MUL, ""
		case '/':
			return l.position, DIV, ""
		// other operators
		case '=':
			return l.position, ASSIGN, ""
		case ';':
			return l.position, SEMI, ""
		// parents
		case '(':
			return l.position, PARENT_LEFT, ""
		case ')':
			return l.position, PARENT_RIGHT, ""
		// sections
		case '{':
			return l.position, SECTION_LEFT, ""
		case '}':
			return l.position, SECTION_RIGHT, ""
		// others
		default:
			if unicode.IsSpace(r) { // skip
				continue // nothing, just skip...
			} else if unicode.IsDigit(r) { // int
				startPos := l.position
				l.back()
				// back and let parseInt check the ident at the first...
				lit := l.parseInt()
				return startPos, INT, lit
			} else if unicode.IsLetter(r) { // ident
				startPos := l.position
				l.back()
				// back and let parseIdent check the ident at the first...
				lit := l.parseIdent()
				if lit == "ret" {
					return startPos, RET, ""
				} else {
					return startPos, IDENT, lit
				}
			} else { // error, unknown
				return l.position, UNKNOWN, string(r)
			}
		}
	}
}

func (l *Lexer) nextLine() {
	l.position.line++
	l.position.column = 0
}

func (l *Lexer) back() {
	if err := l.reader.UnreadRune(); err != nil {
		panic(err)
	}

	l.position.column--
}

// parseInt() scans digit number and returns the literal.
func (l *Lexer) parseInt() string {
	var literal string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// at the end of the int
				return literal
			}
		}

		l.position.column++
		if unicode.IsDigit(r) {
			literal = literal + string(r)
		} else {
			// scanned something not in the integer
			l.back()
			return literal
		}
	}
}

// parseIdent scans ident/variable name and returns the literal.
func (l *Lexer) parseIdent() string {
	var literal string
	for {
		r, _, err := l.reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				// end of the file, so end of the literal ident
				return literal
			}
		}

		l.position.column++
		if unicode.IsLetter(r) {
			// append new letter to the literal temp string
			literal = literal + string(r)
		} else {
			// not a valid letter, go back and so end of the literal ident
			l.back()
			return literal
		}
	}
}
