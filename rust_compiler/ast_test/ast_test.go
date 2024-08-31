package ast_test

import (
	"rust_compiler/ast"
	"rust_compiler/tokens"
	"testing"
)

func TestString(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: tokens.Token{Type: tokens.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: tokens.Token{Type: tokens.IDENT, Literal: "foo"},
					Value: "bar",
				},
				Value: &ast.Identifier{
					Token: tokens.Token{Type: tokens.IDENT, Literal: "foofoo"},
					Value: "barbar",
				},
			},
		},
	}

	if program.String() != "let foo = foofoo;" {
		t.Errorf("program.String() misbehaving; expected 'let foo = foofoo;' and got %q", program.String())
	}
}
