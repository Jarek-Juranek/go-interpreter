package lexer

import (
	"monkey_interpreter/token"
)

type Lexer struct {
	input        string
	position     int  // current position lexer is on to
	readPosition int  // position that will be red next
	ch           byte // current char that is examined
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// Read next char from input and move position pointer further
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, '=')
	case '+':
		tok = newToken(token.ADD, '+')
	case '(':
		tok = newToken(token.LPAREN, '(')
	case ')':
		tok = newToken(token.RPAREN, ')')
	case '{':
		tok = newToken(token.LBRACE, '{')
	case '}':
		tok = newToken(token.RBRACE, '}')
	case ',':
		tok = newToken(token.COMMA, ',')
	case ';':
		tok = newToken(token.SEMICOLON, ';')
	case 0:
		tok = newToken(token.EOF, 0)
	}

	l.readChar()
	return tok
}

// Helper method to create token cleaner
func newToken(tokenType token.TokenType, ch byte) token.Token {
	if ch == 0 {
		return token.Token{Type: tokenType, Literal: ""}
	}
	return token.Token{Type: tokenType, Literal: string(ch)}
}
