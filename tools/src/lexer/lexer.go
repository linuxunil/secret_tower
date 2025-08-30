package lexer

import (
	"secret_compiler/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

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
	case '#':
		tok = newToken(token.H, l.ch)
	case '-':
		tok = newToken(token.SEP, l.ch)
	case '>':
		tok = newToken(token.DIR, l.ch)
	case '[':
		tok = newToken(token.METAL, l.ch)
	case ']':
		tok = newToken(token.METAR, l.ch)
	case '*':
		tok = newToken(token.OBJ, l.ch)
	}
	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
