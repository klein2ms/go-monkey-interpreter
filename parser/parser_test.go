package parser

import (
	"fmt"
	"github.com/klein2ms/go-monkey-interpreter/ast"
	"github.com/klein2ms/go-monkey-interpreter/lexer"
	"testing"
)

func TestLetStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"let x = 5;", "x", 5},
		{"let y = true;", "y", true},
		{"let foobar = y;", "foobar", "y"},
	}

	for _, tt := range tests {
		program := testStatement(t, tt.input, 1)
		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}

		val := stmt.(*ast.LetStatement).Value
		if !testLiteralExpression(t, val, tt.expectedValue) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
		expectedValue      interface{}
	}{
		{"return 5;", "return", 5},
		{"return x;", "return", "x"},
		{"return true;", "return", true},
	}

	for _, tt := range tests {
		program := testStatement(t, tt.input, 1)

		stmt := program.Statements[0]
		returnStmt, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("s not *ast.ReturnStatement. got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("returnStmt.TokenLiteral not 'return', got %q", returnStmt.TokenLiteral())
		}

		val := stmt.(*ast.ReturnStatement).ReturnValue
		if !testLiteralExpression(t, val, tt.expectedValue) {
			return
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

		testIdentifier(t, expStmt.Expression, tt.expectedIdentifier)
	}
}

func TestIntegerLiteralExpression(t *testing.T) {
	input := "5;"

	program := testStatement(t, input, 1)

	tests := []struct {
		expectedValue int64
	}{
		{5},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		expStmt, ok := stmt.(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("exp not *ast.ExpressionStatement. got=%T", stmt)
		}

		testIntegerLiteral(t, expStmt.Expression, tt.expectedValue)
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{"!5;", "!", 5},
		{"-15;", "-", 15},
		{"!true", "!", true},
		{"!false", "!", false},
	}

	for _, tt := range prefixTests {
		program := testStatement(t, tt.input, 1)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s", tt.operator, exp.Operator)
		}

		if !testLiteralExpression(t, exp.Right, tt.value) {
			return
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
		{"true == true", true, "==", true},
		{"true != false", true, "!=", false},
		{"false == false", false, "==", false},
	}

	for _, tt := range infixTests {
		program := testStatement(t, tt.input, 1)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		if !testInfixExpression(t, stmt.Expression, tt.leftValue, tt.operator, tt.rightValue) {
			return
		}
	}
}

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input          string
		expected       string
		statementCount int
	}{
		{"-a * b", "((-a) * b)", 1},
		{"!-a", "(!(-a))", 1},
		{"a + b + c", "((a + b) + c)", 1},
		{"a + b - c", "((a + b) - c)", 1},
		{"a * b * c", "((a * b) * c)", 1},
		{"a * b / c", "((a * b) / c)", 1},
		{"a + b / c", "(a + (b / c))", 1},
		{"a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)", 1},
		{"3 + 4; -5 * 5", "(3 + 4)((-5) * 5)", 2},
		{"5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))", 1},
		{"5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))", 1},
		{"3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))", 1},
		{"true", "true", 1},
		{"false", "false", 1},
		{"3 > 5 == false", "((3 > 5) == false)", 1},
		{"3 < 5 == true", "((3 < 5) == true)", 1},
		{"1 + (2 + 3) + 4", "((1 + (2 + 3)) + 4)", 1},
		{"(5 + 5) * 2", "((5 + 5) * 2)", 1},
		{"2 / (5 + 5)", "(2 / (5 + 5))", 1},
		{"-(5 + 5)", "(-(5 + 5))", 1},
		{"!(true == true)", "(!(true == true))", 1},
		{"a + add(b * c) + d", "((a + add((b * c))) + d)", 1},
		{
			"add(a, b, 1, 2 * 3, 4 + 5, add(6, 7 * 8))",
			"add(a, b, 1, (2 * 3), (4 + 5), add(6, (7 * 8)))",
			1,
		},
		{"add(a + b + c * d / f + g)", "add((((a + b) + ((c * d) / f)) + g))", 1},
	}

	for _, tt := range tests {
		program := testStatement(t, tt.input, tt.statementCount)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func TestBooleanExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"true;", true},
		{"false;", false},
	}

	for _, tt := range tests {
		program := testStatement(t, tt.input, 1)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.BooleanExpression)

		if !ok {
			t.Fatalf("exp is not ast.BooleanExpression. got=%T", stmt)
		}

		if exp.Value != tt.expected {
			t.Fatalf("exp.Value is not %t. got=%t", exp.Value, tt.expected)
		}
	}
}

func TestIfExpression(t *testing.T) {
	tests := []struct {
		input          string
		expected       string
		statementCount int64
	}{
		{`if (x < y) { x }`, ``, 1},
	}

	for _, tt := range tests {
		program := testStatement(t, tt.input, 1)

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.IfExpression)

		if !ok {
			t.Fatalf("exp is not ast.IfExpression. got=%T", stmt)
		}

		if !testInfixExpression(t, exp.Condition, "x", "<", "y") {
			return
		}

		if len(exp.Consequence.Statements) != 1 {
			t.Errorf("consequence is not 1 statement. got=%d\n", len(exp.Consequence.Statements))
		}

		consequence, ok := exp.Consequence.Statements[0].(*ast.ExpressionStatement)

		if !ok {
			t.Fatalf(
				"exp.Consequence.Statements[0] is not ast.ExpressionStatement. got=%T",
				exp.Consequence.Statements[0])
		}

		if !testIdentifier(t, consequence.Expression, "x") {
			return
		}

		if exp.Alternative != nil {
			t.Errorf("exp.Alternative.Statements was not nil. got=%+v", exp.Alternative)
		}
	}
}

func TestFunctionLiteral(t *testing.T) {
	input := `fn(x, y) { x + y; }`

	program := testStatement(t, input, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		fatalNotExpressionStatement(t, program)
	}

	function, ok := stmt.Expression.(*ast.FunctionLiteral)
	if !ok {
		fatalNotSpecificExpressionStatement(t, stmt, ast.FunctionLiteral{})
	}

	if len(function.Parameters) != 2 {
		t.Fatalf("function literal parameters wrong. want 2, got=%d\n", len(function.Parameters))
	}

	testLiteralExpression(t, function.Parameters[0], "x")
	testLiteralExpression(t, function.Parameters[1], "y")

	if len(function.Body.Statements) != 1 {
		t.Fatalf("function.Body.Statments does not have 1 statment. got=%d\n", len(function.Body.Statements))
	}

	body, ok := function.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("function body is not ast.ExpressionStatement. got=%T", function.Body.Statements[0])
	}

	testInfixExpression(t, body.Expression, "x", "+", "y")
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
		statementCount int64
	}{
		{"fn() {};", []string{}, 1},
		{"fn(x) {}", []string{"x"}, 1},
		{"fn(x, y, z) {}", []string{"x", "y", "z"}, 1},
	}

	for _, tt := range tests {
		program := testStatement(t, tt.input, 1)

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		function := stmt.Expression.(*ast.FunctionLiteral)

		if len(function.Parameters) != len(tt.expectedParams) {
			t.Errorf(
				"length parameters wrong. want %d, got=%d\n",
				len(tt.expectedParams),
				len(function.Parameters))
		}

		for i, ident := range tt.expectedParams {
			testLiteralExpression(t, function.Parameters[i], ident)
		}
	}
}

func TestCallExpression(t *testing.T) {
	input := `add(1, 2 * 3, 4 + 5);`

	program := testStatement(t, input, 1)

	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		fatalNotExpressionStatement(t, program)
	}

	exp, ok := stmt.Expression.(*ast.CallExpression)
	if !ok {
		fatalNotSpecificExpressionStatement(t, stmt, ast.CallExpression{})
	}

	if !testIdentifier(t, exp.Function, "add") {
		return
	}

	if len(exp.Arguments) != 3 {
		t.Fatalf("wrong number of arguments. got=%d\n", len(exp.Arguments))
	}

	testLiteralExpression(t, exp.Arguments[0], 1)
	testInfixExpression(t, exp.Arguments[1], 2, "*", 3)
	testInfixExpression(t, exp.Arguments[2], 4, "+", 5)
}

func fatalNotExpressionStatement(t *testing.T, program *ast.Program) {
	t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
}

func fatalNotSpecificExpressionStatement(t *testing.T, stmt *ast.ExpressionStatement, expectedType interface{}) {
	t.Fatalf("exp is not %T. got=%T", expectedType, stmt)
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

func testLiteralExpression(t *testing.T, exp ast.Expression, expected interface{}) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, exp, int64(v))
	case int64:
		return testIntegerLiteral(t, exp, v)
	case string:
		return testIdentifier(t, exp, v)
	case bool:
		return testBooleanLiteral(t, exp, v)
	}
	t.Errorf("type of exp not handled. got=%T", exp)
	return false
}

func testInfixExpression(
	t *testing.T,
	exp ast.Expression,
	left interface{},
	operator string,
	right interface{}) bool {
	opExp, ok := exp.(*ast.InfixExpression)
	if !ok {
		t.Errorf("exp is not ast.InfexExpression. got=%T(%s)", exp, exp)
		return false
	}

	if !testLiteralExpression(t, opExp.Left, left) {
		return false
	}

	if opExp.Operator != operator {
		t.Errorf("exp.Operator is not '%s'. got=%q", operator, opExp.Operator)
		return false
	}

	if !testLiteralExpression(t, opExp.Right, right) {
		return false
	}

	return true
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}

	if integ.Value != value {
		t.Errorf("integ.Value not %d. got=%d", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value, integ.TokenLiteral())
		return false
	}

	return true
}

func testBooleanLiteral(t *testing.T, exp ast.Expression, value bool) bool {
	boolean, ok := exp.(*ast.BooleanExpression)

	if !ok {
		t.Errorf("exp not *asp.Boolean. got=%T", exp)
		return false
	}

	if boolean.Value != value {
		t.Errorf("boolean.Value not %t. got=%t", value, boolean.Value)
		return false
	}

	if boolean.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("boolean.TokenLiteral not %t. got=%s", value, boolean.TokenLiteral())
		return false
	}

	return true
}

func testIdentifier(t *testing.T, exp ast.Expression, value string) bool {
	ident, ok := exp.(*ast.Identifier)
	if !ok {
		t.Errorf("exp not *ast.Identifier. got=%T", exp)
		return false
	}

	if ident.Value != value {
		t.Errorf("ident.Value not %s. got=%s", value, ident.Value)
		return false
	}

	if ident.TokenLiteral() != value {
		t.Errorf("ident.TokenLiteral not %s. got=%s", value, ident.TokenLiteral())
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
