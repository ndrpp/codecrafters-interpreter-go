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

	IDENT = "IDENTIFIER"

	FUN    = "FUN"
	IF     = "IF"
	ELSE   = "ELSE"
	WHILE  = "WHILE"
	FOR    = "FOR"
	RETURN = "RETURN"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	AND    = "AND"
	OR     = "OR"
	PRINT  = "PRINT"
	CLASS  = "CLASS"
	NIL    = "NIL"
	SUPER  = "SUPER"
	THIS   = "THIS"
	VAR    = "VAR"
)

var keywords = map[string]TokenType{
	"fun":    FUN,
	"if":     IF,
	"while":  WHILE,
	"else":   ELSE,
	"return": RETURN,
	"true":   TRUE,
	"false":  FALSE,
	"and":    AND,
	"class":  CLASS,
	"for":    FOR,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"super":  SUPER,
	"this":   THIS,
	"var":    VAR,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
