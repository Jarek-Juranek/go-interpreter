package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywoards = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywoards[ident]; ok {
		return tok
	}

	return IDENT
}

// Recognized possible tokens:
const (
	// SEPCIAL
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// TYPES
	IDENT = "IDENT" // identifier, like x in let x =...
	INT   = "INT"

	// OPERATORS
	ASSIGN   = "="
	ADD      = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// DELIMITERS, CODE STRUCTURE
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// LANGUAGE KEYWOARDS
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
