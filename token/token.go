package token

type TokenType string

type Token struct {
	Type    TokenType
	Text    string
	Literal string
	Line    uint16
}

const (
	EOF        = "EOF"
	UNEXPECTED = "UNEXPECTED"

	LPAREN = "LEFT_PAREN"
	RPAREN = "RIGHT_PAREN"
	LBRACE = "LEFT_BRACE"
	RBRACE = "RIGHT_BRACE"

	PLUS      = "PLUS"
	MINUS     = "MINUS"
	ASTERISK  = "STAR"
	COMMA     = "COMMA"
	DOT       = "DOT"
	SEMICOLON = "SEMICOLON"
)
