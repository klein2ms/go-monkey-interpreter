package lexer

import (
	"github.com/klein2ms/go-monkey-interpreter/token"
)

// Lexer represents a type used to lex a string input
type Lexer struct {
	input        string
	position     int  //current position in input (points to current char)
	readPosition int  //current reading positon in input (after current char)
	ch           byte //current char under examination
}

// New initializes a new Lexer
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken consumes the next char from the input for a given Lexer and returns a new Token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.read(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}

		if isDigit(l.ch) {
			tok.Literal = l.read(isDigit)
			tok.Type = token.INT
			return tok
		}

		tok = newToken(token.ILLEGAL, l.ch)
	}
	l.readChar()
	return tok
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

func (l *Lexer) read(a func(ch byte) bool) string {
	position := l.position
	for a(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// Helps inititialize tokens
func newToken(tokenType token.Type, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
