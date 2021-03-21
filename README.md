# mylex: go-lexer-token-simple

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
./go-lexer-token-simple
```

or using go get:

```bash
go get github.com/BaseMax/go-lexer-token-simple
cd /path/to/go-lexer-token-simple
go build
./go-lexer-token-simple
```

## Output

```
5:1	IDENT	int
5:5	IDENT	main
5:9	(	
5:10	)	
5:12	{	
6:2	IDENT	a
6:4	=	
6:6	INT	5
6:7	;	
7:2	IDENT	b
7:4	=	
7:6	IDENT	a
7:8	+	
7:10	INT	6
7:11	;	
8:2	IDENT	c
8:4	+	
8:6	INT	123
8:9	;	
9:2	INT	5
9:3	+	
9:4	INT	12
9:6	;	
10:2	RET	
10:6	IDENT	c
10:7	;	
11:1	}	
```

## TODO

- Support \r\n
- Support floating numbers

## My references

- How `UnreadRune` works? https://golang.org/pkg/bufio/#Reader.UnreadRune
- Anko: a perfect go-script language made with Go: https://github.com/mattn/anko
- Lark: a old but funny lexer-interpreter made with Java: https://github.com/munificent/lark/blob/master/src/com/stuffwithstuff/lark/Lexer.java
- How `iota` works in golang? https://yourbasic.org/golang/iota/
- How write a simple go lexer? https://aaronraff.dev/writing-a-lexer-in-go
- Why use `iota` in golang? https://medium.com/swlh/iota-create-effective-constants-in-golang-b399f94aac31

© Copyright Max Base
