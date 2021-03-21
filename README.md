# go-lexer-token-simple

Simple Go lexer: Lex own syntax and read it's from file.

## Features

- Tokenizer
- Position analyzer (line and offset)
- Ident
- Int
- Comments # ...
- Operators (+ - * / =)

## Example / Syntax

```
# Max Base
# simple lexer written in GoLang
# 2021-03-21

int main() {
	a = 5;
	b = a + 6;
	c + 123;
	5+12;
	ret c;
}
```

## Using

```bash
git clone https://github.com/BaseMax/go-lexer-token-simple
cd go-lexer-token-simple
go build
./lex
```

or using go get:

```bash
go get github.com/BaseMax/go-lexer-token-simple
cd /path/to/go-lexer-token-simple
go build
./lex
```

## My references

- How `UnreadRune` works? https://golang.org/pkg/bufio/#Reader.UnreadRune
- Anko: a perfect go-script language made with Go: https://github.com/mattn/anko
- Lark: a old but funny lexer-interpreter with with Java: https://github.com/munificent/lark/blob/master/src/com/stuffwithstuff/lark/Lexer.java
- How `iota` works in golang? https://yourbasic.org/golang/iota/
- How write a simple go lexer? https://aaronraff.dev/writing-a-lexer-in-go
- Why use `iota` in golang? https://medium.com/swlh/iota-create-effective-constants-in-golang-b399f94aac31

© Copyright Max Base
