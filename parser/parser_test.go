package parser

import (
	"github.com/klein2ms/go-monkey-interpreter/ast"
	"github.com/klein2ms/go-monkey-interpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	program := testStatement(t, input, 3)

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	`

	program := testStatement(t, input, 3)

	for _, stmt := range program.Statements {
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("s not *ast.ReturnStatement. got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}
	}
}

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	program := testStatement(t, input, 1)

	tests := []struct {
		expectedIdentifier string
	}{
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		expStmt, ok := stmt.(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("exp not *ast.ExpressionStatement. got=%T", stmt)
		}

		ident, ok := expStmt.Expression.(*ast.Identifier)
		if !ok {
			t.Fatalf("exp not *ast.Identifier. got=%T", expStmt.Expression)
		}

		if ident.Value != tt.expectedIdentifier {
			t.Errorf("ident.Value not %s. got=%s", tt.expectedIdentifier, ident.Value)
		}

		if ident.TokenLiteral() != tt.expectedIdentifier {
			t.Errorf("ident.TokenLiteral not %s. got=%s", tt.expectedIdentifier, ident.TokenLiteral())
		}
	}
}

func testStatement(t *testing.T, input string, statementCount int) *ast.Program {
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != statementCount {
		t.Fatalf("program.Statements does not contain %d statements. got=%d", statementCount, len(program.Statements))
	}

	return program
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.TokenLiteral() not '%s'.. got=%s", name, letStmt.Name.TokenLiteral())
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser had %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
