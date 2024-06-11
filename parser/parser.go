package parser

import (
	"hal/ast"
	"hal/lexer"
	"hal/token"
)

// Parser 3 fields: l, currToken and peekToken
// l is a pointer to an instance of the lexer, on which we repeatedly call NextToken to get the next token of the input
// currToken and peekToken are 2 pointers to the current token and the next token, similar to the pointers the lexer
// has, but it's to tokens instead of characters
// we need to look at the current token to decide what to do next, and we need to look at peekToken if currToken
// doesn't give enough information.
type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

// New create a new parser
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// read 2 tokens so currToken and peekToken are both set
	p.nextToken()
	p.nextToken()

	return p
}

// nextToken advances both currToken and peekToken
func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
