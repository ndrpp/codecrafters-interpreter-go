package main

import (
	"fmt"
	"os"

	"github.com/codecrafters-io/interpreter-starter-go/token"
)

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

	hadScanErrors := false
	line := 0
	if len(fileContents) > 0 {
	loop:
		for i := 0; i < len(fileContents); i++ {
			var tok token.Token
			v := fileContents[i]

			switch v {
			case ' ':
				continue loop
			case '\r':
				continue loop
			case '\t':
				continue loop
			case '\n':
				line++
				continue loop
			case '(':
				tok = newToken(token.LPAREN, v)
			case ')':
				tok = newToken(token.RPAREN, v)
			case '}':
				tok = newToken(token.RBRACE, v)
			case '{':
				tok = newToken(token.LBRACE, v)
			case '+':
				tok = newToken(token.PLUS, v)
			case '-':
				tok = newToken(token.MINUS, v)
			case '*':
				tok = newToken(token.ASTERISK, v)
			case ',':
				tok = newToken(token.COMMA, v)
			case '.':
				tok = newToken(token.DOT, v)
			case ';':
				tok = newToken(token.SEMICOLON, v)
			case '=':
				next := peekNextToken(i, fileContents)
				if next == '=' {
					tok = token.Token{Type: token.EQ, Literal: string(v) + string(next)}
					i++
				} else {
					tok = newToken(token.ASSIGN, v)
				}
			case '!':
				next := peekNextToken(i, fileContents)
				if next == '=' {
					tok = token.Token{Type: token.NOT_EQ, Literal: string(v) + string(next)}
					i++
				} else {
					tok = newToken(token.BANG, v)
				}
			case '>':
				next := peekNextToken(i, fileContents)
				if next == '=' {
					tok = token.Token{Type: token.GT_OR_EQ, Literal: string(v) + string(next)}
					i++
				} else {
					tok = newToken(token.GT, v)
				}
			case '<':
				next := peekNextToken(i, fileContents)
				if next == '=' {
					tok = token.Token{Type: token.LT_OR_EQ, Literal: string(v) + string(next)}
					i++
				} else {
					tok = newToken(token.LT, v)
				}
			case '/':
				next := peekNextToken(i, fileContents)
				if next == '/' {
					i++
					for {
						ch := peekNextToken(i, fileContents)
						if ch == '\n' || i+1 == len(fileContents) {
							break
						} else {
							i++
						}
					}
					continue loop
				} else {
					tok = newToken(token.SLASH, v)
				}
			case '"':
				init := i
				for {
					ch := peekNextToken(i, fileContents)
					if ch == '\n' || ch == '"' || ch == 0 {
						break
					} else {
						i++
					}
				}
				if i+1 == len(fileContents) {
					fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", line+1)
					hadScanErrors = true
					continue loop
				}
				tok.Type = token.STRING
				tok.Literal = string(fileContents[init : i+2])
				tok.Text = string(fileContents[init+1 : i+1])
				fmt.Printf("%s %s %s\n", tok.Type, tok.Literal, tok.Text)
				i++
				continue loop

			default:
				fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %s\n", line+1, string(v))
				hadScanErrors = true
				continue loop
			}

			fmt.Printf("%s %s null\n", tok.Type, tok.Literal)
		}
	}
	fmt.Println("EOF  null")

	if hadScanErrors {
		os.Exit(65)
	}
}

func peekNextToken(i int, fileContents []byte) byte {
	if i+1 < len(fileContents) {
		return fileContents[i+1]
	}
	return 0
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
