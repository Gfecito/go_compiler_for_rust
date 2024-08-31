package tokens

type TokenType string

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers and types
	IDENT TokenType = "IDENT"
	INT   TokenType = "INT"

	// Operators
	ASSIGN TokenType = "="
	PLUS   TokenType = "+"

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
	RETURN   TokenType = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}
