package ast

// Abstract Syntax Tree (AST)

import (
	"rust_compiler/tokens"
)

type Node interface {
	TokenLiteral() string
}

type Expression interface {
	Node
	expressionNode()
}

type Statement interface {
	Node
	statementNode()
	String() string
}

// Root of the AST.
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (p *Program) String() string {
	toString := ""
	programLength := len(p.Statements)
	for index, statement := range p.Statements {
		toString += statement.String()
		// Add spacing between tokens, unless EOF.
		if index < (programLength - 1) {
			toString += "\n"
		}
	}
	return toString
}

// Compile-time check that Program implements Node
var _ Node = (*Program)(nil)

type Identifier struct {
	Token tokens.Token
	Value string
}

func (identifier *Identifier) expressionNode()      {}
func (identifier *Identifier) TokenLiteral() string { return identifier.Token.Literal }

// Compile-time check that Identifier implements Expression
var _ Expression = (*Identifier)(nil)

// IntegerLiteral node
type IntegerLiteral struct {
	Token tokens.Token // the integer token
	Value int64
}

func (il *IntegerLiteral) expressionNode()      {}
func (il *IntegerLiteral) TokenLiteral() string { return il.Token.Literal }

// Compile-time check that IntegerLiteral implements Expression
var _ Expression = (*IntegerLiteral)(nil)

type LetStatement struct {
	Token tokens.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode()       {}
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }
func (ls *LetStatement) String() string {
	return ls.TokenLiteral() + " " + ls.Name.TokenLiteral() + " = " + ls.Value.TokenLiteral() + ";"
}

// Compile-time check that LetStatement implements Statement
var _ Statement = (*LetStatement)(nil)

type FunctionDeclaration struct {
	Token     tokens.Token
	Name      *Identifier
	Body      Expression
	Arguments Expression
}

func (functionDeclaration *FunctionDeclaration) statementNode() {}
func (functionDeclaration *FunctionDeclaration) TokenLiteral() string {
	return functionDeclaration.Token.Literal
}
func (functionDeclaration *FunctionDeclaration) String() string {
	// TODO
	return ""
}

// Compile-time check that FunctionDeclaration implements Statement
var _ Statement = (*FunctionDeclaration)(nil)
