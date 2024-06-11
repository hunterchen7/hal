package ast

import "hal/token"

type Node interface {
	TokenLiteral() string // used only for debugging purposes
}

// AST needs 2 different types of nodes: expressions and statements.
// they contain statementNode() and expressionNode() which are not strictly necessary
// but help us by guiding the Go compiler and throwing errors when we use a statement
// where an expression should be, and vice versa.

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

// the Program node is the root node of every AST produced by the parser
// every valid program is a series of statements, which are contained in Program.Statements,
// which is just a slice of AST nodes that implement the Statement interface

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (ls *LetStatement) StatementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

// 3 fields, 1 for identifier, 1 for expression that produces the value, and one for the token

type LetStatement struct {
	Token token.Token
	Name  *Identifier // holds the identifier of the binding
	Value Expression  // the value of the expression that produces the value
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
