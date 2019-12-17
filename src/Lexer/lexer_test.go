//lexer/lexer_test2.go

package lexer

import (
	"testing"
	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := '=+(){},;'

	tests := []struct {
		expectedType 	token.TokenType
		expectedLiteral string		
	}{
			{token.ASSIGN,    "="},
			{token.PLUS,      "+"},
			{token.LPAREN,    "("},
			{token.RPAREN,    ")"},
			{token.LBRACE,    "{"},
			{token.COMMA,     ","},
			{token.SEMICOLON, ";"},
			{token.EOF,		   ""},
	}

	l := New(input)

	for i, tt:= range tests {
		tok := l.NextToken()

		if tok.type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}