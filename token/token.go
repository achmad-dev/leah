package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	IDENT   = "IDENT"
	INT     = "INT" // NUMBER

	// OPERATOR
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	REMAINDER = "%"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"

	EQUAL    = "=="
	NOTEQUAL = "!="

	GT = ">"
	LT = "<"

	// DELIMITER
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// KEYWORDS
	FUNCTION = "FUNCTION"
	VAR      = "VAR"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	RETURN   = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"var":    VAR,
	"if":     IF,
	"else":   ELSE,
	"true":   FALSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if token, ok := keywords[ident]; ok {
		return token
	}
	return IDENT
}
