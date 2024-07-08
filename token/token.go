package token

type TokenType string

type Token struct {
	Type    TokenType
	Text    string
	Literal string
	Line    uint16
}

const (
	EOF = "EOF"

	LPAREN = "LEFT_PAREN"
	RPAREN = "RIGHT_PAREN"
)
