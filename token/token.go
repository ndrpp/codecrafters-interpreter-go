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

	STRING = "STRING"
	NUMBER = "NUMBER"

	LPAREN = "LEFT_PAREN"
	RPAREN = "RIGHT_PAREN"
	LBRACE = "LEFT_BRACE"
	RBRACE = "RIGHT_BRACE"

	PLUS      = "PLUS"
	MINUS     = "MINUS"
	ASTERISK  = "STAR"
	SLASH     = "SLASH"
	COMMA     = "COMMA"
	DOT       = "DOT"
	SEMICOLON = "SEMICOLON"
	ASSIGN    = "EQUAL"
	BANG      = "BANG"

	EQ     = "EQUAL_EQUAL"
	NOT_EQ = "BANG_EQUAL"

	LT       = "LESS"
	GT       = "GREATER"
	LT_OR_EQ = "LESS_EQUAL"
	GT_OR_EQ = "GREATER_EQUAL"
)
