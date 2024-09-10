package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
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
	ASSIGN = "="
	ADD    = "+"

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
