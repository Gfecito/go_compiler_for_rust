package semantic

import (
	"fmt"
	"rust_compiler/ast"
)

type BorrowChecker struct {
	ownership map[string]string
	errors    []string
}

func NewBorrowChecker() *BorrowChecker {
	return &BorrowChecker{ownership: make(map[string]string), errors: []string{}}
}

func (bc *BorrowChecker) Check(program *ast.Program) {
	for _, stmt := range program.Statements {
		bc.checkStatement(stmt)
	}
}

func (bc *BorrowChecker) checkStatement(stmt ast.Statement) {
	switch stmt := stmt.(type) {
	case *ast.LetStatement:
		bc.checkLetStatement(stmt)
	default:
	}
}

func (bc *BorrowChecker) checkLetStatement(stmt *ast.LetStatement) {
	if stmt.Value == nil {
		return
	}
	if _, ok := stmt.Value.(*ast.Identifier); ok {
		if bc.ownership[stmt.Name.Value] == "mut_borrowed" {
			bc.errors = append(bc.errors, fmt.Sprintf("cannot borrow %s as immutable because it is already borrowed as mutable", stmt.Name.Value))
		}
		bc.ownership[stmt.Name.Value] = "borrowed"
	}
}

func (bc *BorrowChecker) Errors() []string {
	return bc.errors
}
