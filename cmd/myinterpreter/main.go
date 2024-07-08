package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/token"
)

type Scanner struct {
	start   uint16
	current uint16
	line    uint16
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}

	command := os.Args[1]

	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}

	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		for _, v := range fileContents {
			token := scanToken(v)
			fmt.Printf("%s %s null\n", token.Type, token.Literal)
		}
	}
	fmt.Println("EOF  null")
}

func scanToken(t byte) token.Token {
	var tok token.Token

	switch t {
	case '(':
		tok = newToken(token.LPAREN, t)
	case ')':
		tok = newToken(token.RPAREN, t)
	case '}':
		tok = newToken(token.RBRACE, t)
	case '{':
		tok = newToken(token.LBRACE, t)
	case '+':
		tok = newToken(token.PLUS, t)
	case '-':
		tok = newToken(token.MINUS, t)
	case '*':
		tok = newToken(token.ASTERISK, t)
	case ',':
		tok = newToken(token.COMMA, t)
	case '.':
		tok = newToken(token.DOT, t)
	case ';':
		tok = newToken(token.SEMICOLON, t)

	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
