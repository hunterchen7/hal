package parser

import (
	"fmt"
	"hal/ast"
	"hal/lexer"
	"hal/token"
)

// Parser has 3 fields: l, currToken and peekToken
// l is a pointer to an instance of the lexer, on which we repeatedly call NextToken to get the next token of the input
// currToken and peekToken are 2 pointers to the current token and the next token, similar to the pointers the lexer
// has, but it's to tokens instead of characters
// we need to look at the current token to decide what to do next, and we need to look at peekToken if currToken
// doesn't give enough information.
type Parser struct {
	l *lexer.Lexer

	// errors is just a slice of strings, it gets initialized in New
	// the helper peekError is used to add an error to errors when the type of peekToken doesn't match expectations
	errors []string

	currToken token.Token
	peekToken token.Token
}

// New create a new parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}

	// read 2 tokens so currToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}

// nextToken advances both currToken and peekToken
func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// ParseProgram recursive descent parser
// construct the root node of the AST
// build the child nodes, i.e. the statements by calling other functions that know which AST node to construct
// based on the current token (let, return, etc.)
// these functions call each other recursively
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.currToken.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}

	return program
}

// TODO: add more branches
func (p *Parser) parseStatement() ast.Statement {
	switch p.currToken.Type {
	case token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// construct an *ast.LetStatement node with the token it's currently sitting on
	statement := &ast.LetStatement{Token: p.currToken}

	// it first expects an IDENT token
	if p.peekToken.Type != token.IDENT {
		return nil
	}
	p.nextToken()

	statement.Name = &ast.Identifier{
		Token: p.currToken,
		Value: p.currToken.Literal,
	}

	// then it expects an ASSIGN token i.e. =
	if p.peekToken.Type != token.ASSIGN {
		return nil
	}
	p.nextToken()

	// TODO: implement this
	// currently just skips expressions until a semicolon is encountered
	for p.currToken.Type != token.SEMICOLON {
		p.nextToken()
	}

	return statement
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekToken.Type == t {
		p.nextToken()
		return true
	}
	p.peekError(t)
	return false
}
