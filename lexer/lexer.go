package lexer

import "hal/token"

type Lexer struct {
	input        string
	position     int  // current position in inputs (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// if at end of input, sets l.ch to 0, which is ASCII for "nul"
	// signifies "we haven't read anything yet" or "end of file"
	if l.readPosition >= len(l.input) {
		l.ch = 0
		// otherwise read next character
	} else {
		l.ch = l.input[l.readPosition]
	}
	// position updated to the just used l.readPosition
	l.position = l.readPosition
	// readPosition is incremented by 1 to the next position
	l.readPosition += 1
}

// serves a similar purpose to readChar except it doesn't increment
// position and readPosition. We only want to "peek" the next character,
// and not move around to it.
// Most lexers and parsers have a "peek" function that looks ahead and usually
// only returns the next character. The difficulty of parsing different languages
// often comes down to how far ahead you have to peek ahead (or backwards)
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// this function is used specifically for when a token is 2 characters, it is meant as replacements
// of the if statements found on page 24.
// I don't specify that it must be equal to '=' in its usage, since there may be other potential
// token endings, like "++" or something similar, which may be implemented in the future.
func (l *Lexer) doubleToken(tokenType token.TokenType) token.Token {
	ch := l.ch
	l.readChar()
	return token.Token{
		Type:    tokenType,
		Literal: string(ch) + string(l.ch),
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	switch l.ch {
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			// readChar() is called repeatedly when calling readIdentifier() so readPosition
			// and position are advanced past the last char of the current identifier.
			// therefore we don't need to call readChar() again after this switch
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	case '=':
		if l.peekChar() == '=' {
			tok = l.doubleToken(token.EQ)
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		if l.peekChar() == '=' {
			tok = l.doubleToken(token.PLUSEQ)
		} else {
			tok = newToken(token.PLUS, l.ch)
		}
	case '-':
		if l.peekChar() == '=' {
			tok = l.doubleToken(token.MINUSEQ)
		} else {
			tok = newToken(token.MINUS, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			tok = l.doubleToken(token.NEQ)
		} else {
			tok = newToken(token.EXCLAIM, l.ch)
		}
	case '*':
		if l.peekChar() == '=' {
			tok = l.doubleToken(token.MULTEQ)
		} else {
			tok = newToken(token.MULTIPLY, l.ch)
		}
	case '/':
		if l.peekChar() == '=' {
			tok = l.doubleToken(token.DIVEQ)
		} else {
			tok = newToken(token.DIVIDE, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			tok = l.doubleToken(token.LEQ)
		} else {
			tok = newToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			tok = l.doubleToken(token.GEQ)
		} else {
			tok = newToken(token.GT, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// a to z or A to Z or underscore
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
