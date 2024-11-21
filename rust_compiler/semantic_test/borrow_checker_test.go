package semantic_test

import (
	"rust_compiler/ast"
	"rust_compiler/semantic"
	"rust_compiler/tokens"
	"testing"
)

func TestBorrowChecker(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: tokens.Token{Type: tokens.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: tokens.Token{Type: tokens.IDENT, Literal: "x"},
					Value: "5",
				},
				Value: &ast.Identifier{
					Token: tokens.Token{Type: tokens.IDENT, Literal: "x"},
					Value: "y",
				},
			},
		},
	}

	borrowChecer := semantic.NewBorrowChecker()
	borrowChecer.Check(program)

	errors := borrowChecer.Errors()
	if len(errors) != 1 {
		t.Errorf("expected 1 error but got %d", len(errors))
	}

	expectedError := "cannot borrow x as immutable because it is already borrowed as mutable"
	if errors[0] != expectedError {
		t.Errorf("expected error %q but got %q", expectedError, errors[0])
	}
}

func TestBorrowCheckerNoErrors(t *testing.T) {
	program := &ast.Program{
		Statements: []ast.Statement{
			&ast.LetStatement{
				Token: tokens.Token{Type: tokens.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: tokens.Token{Type: tokens.IDENT, Literal: "x"},
					Value: "5",
				},
				Value: &ast.IntegerLiteral{
					Token: tokens.Token{Type: tokens.INT, Literal: "5"},
					Value: 5,
				},
			},
			&ast.LetStatement{
				Token: tokens.Token{Type: tokens.LET, Literal: "let"},
				Name: &ast.Identifier{
					Token: tokens.Token{Type: tokens.IDENT, Literal: "y"},
					Value: "5",
				},
				Value: &ast.IntegerLiteral{
					Token: tokens.Token{Type: tokens.INT, Literal: "5"},
					Value: 5,
				},
			},
		},
	}

	borrowChecer := semantic.NewBorrowChecker()
	borrowChecer.Check(program)

	errors := borrowChecer.Errors()
	if len(errors) != 0 {
		t.Errorf("expected 0 errors but got %d", len(errors))
	}
}
