package ast

import (
	"github.com/klein2ms/go-monkey-interpreter/token"
	"testing"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, literal: "let"},
				Name: & Identifier{
					Token: token.Token{Type: token.IDENT, literal: "myVar"},
					Value: "myVar",
				},
			},
		},
	}

	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}