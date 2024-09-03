package ast

import "rust_compiler/tokens"

type Type interface {
	Node
	typeNode()
}

type IdentifierType struct {
	Token tokens.Token
	Value string
}

func (idType *IdentifierType) typeNode()            {}
func (idType *IdentifierType) TokenLiteral() string { return idType.Token.Literal }
