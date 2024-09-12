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

	l.skipWhiteSpace()

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
	case '!':
		tok = newToken(token.BANG, '!')
	case '-':
		tok = newToken(token.MINUS, '-')
	case '*':
		tok = newToken(token.ASTERISK, '*')
	case '/':
		tok = newToken(token.SLASH, '/')
	case '<':
		tok = newToken(token.LT, '<')
	case '>':
		tok = newToken(token.GT, '>')
	case 0:
		tok = newToken(token.EOF, 0)
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[startPosition:l.position]
}

func (l *Lexer) readNumber() string {
	startPosition := l.position

	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[startPosition:l.position]
}

func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Helper methods

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	if ch == 0 {
		return token.Token{Type: tokenType, Literal: ""}
	}
	return token.Token{Type: tokenType, Literal: string(ch)}
}
