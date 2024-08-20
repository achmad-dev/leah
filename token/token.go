package token

type tokenType string

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

	SAME    = "=="
	NOTSAME = "!="

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

type token struct {
	Type    tokenType
	Literal string
}

var keywords = map[string]tokenType{
	"fn":     FUNCTION,
	"var":    VAR,
	"if":     IF,
	"else":   ELSE,
	"true":   FALSE,
	"return": RETURN,
}

func LookupIdent(ident string) tokenType {
	if token, ok := keywords[ident]; ok {
		return token
	}
	return IDENT
}
