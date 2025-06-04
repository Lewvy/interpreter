package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	LineNo  int
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"false":  FALSE,
	"true":   TRUE,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	MINUS     = "-"
	SLASH     = "/"
	BANG      = "!"
	ASTERISK  = "*"
	LT        = "<"
	GT        = ">"
	EQ        = "=="
	NOT_EQ    = "!="

	FUNCTION = "FUNCTION"
	LET      = "LET"
	FALSE    = "FALSE"
	TRUE     = "TRUE"
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
)
