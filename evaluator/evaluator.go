package evaluator

import (
	"github.com/klein2ms/go-monkey-interpreter/ast"
	"github.com/klein2ms/go-monkey-interpreter/object"
)

var (
	// NULL represents the shared null object
	NULL = &object.Null{}
	// TRUE represents the shared boolean object for true
	TRUE = &object.Boolean{Value: true}
	// FALSE represents the shared boolean object for false
	FALSE = &object.Boolean{Value: false}
)

// Eval evaluates a given ast node return the representative object type and boxed value
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalStatements(node.Statements)

	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.BooleanExpression:
		return nativeBoolToBooleanObject(node.Value)

	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	}

	return nil
}

func evalStatements(statements []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range statements {
		result = Eval(statement)
	}

	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}

	return FALSE
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)
	default:
		return NULL
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE, NULL:
		return FALSE
	case FALSE:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER {
		return NULL
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}
