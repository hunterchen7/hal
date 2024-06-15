package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// Operators
	PLUS     = "+"
	MINUS    = "-"
	EXCLAIM  = "!"
	MULTIPLY = "*"
	DIVIDE   = "/"

	// Assignments
	ASSIGN  = "="
	PLUSEQ  = "+="
	MINUSEQ = "-="
	MULTEQ  = "*="
	DIVEQ   = "/="

	// Comparators
	LT  = "<"
	GT  = ">"
	LEQ = "<="
	GEQ = ">="
	EQ  = "=="
	NEQ = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fun":    FUNCTION,
	"let":    LET,
	"return": RETURN,
}

// LookupIdent checks the keywords table to see if a given identifier is a keyword
func LookupIdent(ident string) TokenType {
	// returns token if it is
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	// returns IDENT if isn't, which is the token type for all user identified identifiers
	return IDENT
}
