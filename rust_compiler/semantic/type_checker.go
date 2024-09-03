package semantic

import (
	"fmt"
	"rust_compiler/ast"
	"rust_compiler/symbol"
)

type TypeChecker struct {
	symbolTable *symbol.SymbolTable
	errors      []string
}

func NewTypeChecker(symbolTable *symbol.SymbolTable) *TypeChecker {
	return &TypeChecker{symbolTable: symbolTable, errors: []string{}}
}

func (tc *TypeChecker) Check(program *ast.Program) {
	for _, stmt := range program.Statements {
		tc.checkStatement(stmt)
	}
}

func (tc *TypeChecker) checkStatement(stmt ast.Statement) {
	switch stmt := stmt.(type) {
	case *ast.LetStatement:
		tc.checkLetStatement(stmt)
	default:
		//
	}
}

func (tc *TypeChecker) checkLetStatement(letStmt *ast.LetStatement) {
	// The 'foo' in 'let foo = bar'
	variableName := letStmt.Name.Value
	expectedType := tc.symbolTable.Lookup(variableName)
	// Didn't find type for this variable!
	if expectedType == nil {
		tc.errors = append(tc.errors, fmt.Sprintf("undefined variable: %s", variableName))
		return
	}
	// The 'bar' in 'let foo = bar'
	actualType := tc.evaluateExpression(letStmt.Value)
	if expectedType != actualType {
		tc.errors = append(tc.errors, fmt.Sprintf("type mismatch: expected %s but got %s", expectedType, actualType))
	}
}

func (tc *TypeChecker) evaluateExpression(expr ast.Expression) symbol.Symbol {
	switch expr.(type) {
	case *ast.IntegerLiteral:
		return symbol.IntSymbol
	default:
		return symbol.UndefinedSymbol
	}
}
