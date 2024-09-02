package parser

import (
	"rust_compiler/ast"
	"rust_compiler/lexer"
	"rust_compiler/tokens"
)

type Parser struct {
	l         *lexer.Lexer
	curToken  tokens.Token
	peekToken tokens.Token
	Errors    []string
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		Errors: []string{},
	}

	// Peek at the first Token
	p.nextToken()
	// First Token will now be current, and we're peeking at the second one.
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case tokens.LET:
		return p.parseLetStatement()
	// case tokens.RETURN:
	// 	return p.parseReturnStatement()
	default:
		return nil
		// return p.parseExpressionStatement()
	}
}

func (p *Parser) parseReturnStatement() *ast.LetStatement {
	// stmt is an abbreviation for statement.
	stmt := &ast.LetStatement{Token: p.curToken}

	return stmt
}

func (p *Parser) parseExpressionStatement() *ast.LetStatement {
	// stmt is an abbreviation for statement.
	stmt := &ast.LetStatement{Token: p.curToken}

	return stmt
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// stmt is an abbreviation for statement.
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectPeek(tokens.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectPeek(tokens.ASSIGN) {
		return nil
	}

	p.nextToken()

	// stmt.Value = *p.parseExpressionStatement(LOWEST)

	if p.expectPeek(tokens.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) expectPeek(expectedType tokens.TokenType) bool {
	return p.peekToken.Type == expectedType
}

// Returns the root node of the AST: the program!
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != tokens.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}

	return program
}
