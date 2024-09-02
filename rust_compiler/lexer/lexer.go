package lexer

import (
	"rust_compiler/tokens"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// End of input.
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() tokens.Token {
	var tokenType tokens.TokenType
	var tok tokens.Token

	l.skipWhitespace()

	switch l.ch {
	// Special characters
	case '=':
		tokenType = tokens.ASSIGN
	case '+':
		tokenType = tokens.PLUS
	case ',':
		tokenType = tokens.COMMA
	case ';':
		tokenType = tokens.SEMICOLON
	case '(':
		tokenType = tokens.LPAREN
	case ')':
		tokenType = tokens.RPAREN
	case '{':
		tokenType = tokens.LBRACE
	case '}':
		tokenType = tokens.RBRACE
	case 0:
		tokenType = tokens.EOF
	// Not a special character
	default:
		if isLegalIdentifierChar(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = l.keyWordMatch(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = tokens.INT
			return tok
		} else {
			// Unrecognized token.
			tokenType = tokens.ILLEGAL
		}
	}

	tok = newToken(tokenType, l.ch)

	// We get here on special tokens; we do an early return otherwise.
	// We need to move our reader to the next char.
	// Not doing so would loop infinitely on special characters.
	l.readChar()
	return tok
}

// Looks up an identifier to see if its a reserved keyword.
func (l *Lexer) keyWordMatch(identifier string) tokens.TokenType {
	switch identifier {
	case "let":
		return tokens.LET
	case "fn":
		return tokens.FUNCTION
	case "return":
		return tokens.RETURN
	default:
		return tokens.IDENT
	}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType tokens.TokenType, ch byte) tokens.Token {
	return tokens.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	startingPosition := l.position
	// Read until end of identifier.
	for isLegalIdentifierChar(l.ch) {
		l.readChar()
	}
	// Identifier will be segment between start and end positions.
	return l.input[startingPosition:l.position]
}

func isDigit(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isLegalIdentifierChar(ch byte) bool {
	return isLetter(ch) || ch == '_'
}

func isLetter(ch byte) bool {
	isLowerCase := ch >= 'a' && ch <= 'z'
	isUpperCase := ch >= 'A' && ch <= 'Z'
	return isLowerCase || isUpperCase
}

func (l *Lexer) readNumber() string {
	startingPosition := l.position
	// Read until end of identifier.
	for isDigit(l.ch) {
		l.readChar()
	}
	// Identifier will be segment between start and end positions.
	return l.input[startingPosition:l.position]
}
