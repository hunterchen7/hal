package lexer

import (
	"hal/token"
	"testing"
)

func runTestNextToken(t *testing.T, input string, output []struct {
	expectedType    token.TokenType
	expectedLiteral string
}) {
	l := New(input)

	for i, tt := range output {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("Tests[%d] - tokentype wrong. Expected %q, got %q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("Tests[%d] - literal wrong. Expected %q, got %q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}

func TestNextToken1(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runTestNextToken(t, input, tests)
}

func TestNextToken2(t *testing.T) {
	input := `
let three = 3;
let four = 4;

let add = fn(x, y) {
	x + y;
};

let result = add(three, four);
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "three"},
		{token.ASSIGN, "="},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "four"},
		{token.ASSIGN, "="},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "three"},
		{token.COMMA, ","},
		{token.IDENT, "four"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	runTestNextToken(t, input, tests)
}

func TestNextToken3_extraSymbols(t *testing.T) {
	input := `
let num = 12 / 3;
4 * 3;
4 < 3 > 4534 / 3 ! 34
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "num"},
		{token.ASSIGN, "="},
		{token.INT, "12"},
		{token.DIVIDE, "/"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.INT, "4"},
		{token.MULTIPLY, "*"},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.INT, "4"},
		{token.LT, "<"},
		{token.INT, "3"},
		{token.GT, ">"},
		{token.INT, "4534"},
		{token.DIVIDE, "/"},
		{token.INT, "3"},
		{token.EXCLAIM, "!"},
		{token.INT, "34"},
	}

	runTestNextToken(t, input, tests)
}

func TestNextToken4_doubleTokens(t *testing.T) {
	input := `!=*!= =+(){},; <= >= += -= *= /= < > == let`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.NEQ, "!="},
		{token.MULTIPLY, "*"},
		{token.NEQ, "!="},
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.LEQ, "<="},
		{token.GEQ, ">="},
		{token.PLUSEQ, "+="},
		{token.MINUSEQ, "-="},
		{token.MULTEQ, "*="},
		{token.DIVEQ, "/="},
		{token.LT, "<"},
		{token.GT, ">"},
		{token.EQ, "=="},
		{token.LET, "let"},
		{token.EOF, ""},
	}

	runTestNextToken(t, input, tests)
}
