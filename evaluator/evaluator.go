package evaluator

import (
	"github.com/klein2ms/go-monkey-interpreter/object"
	"github.com/klein2ms/go-monkey-interpreter/ast"
)

var (
	NULL = &object.Null{}
	TRUE =  &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

//Eval function for object.Object
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

//statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

//booleans
	case *ast.Boolean:
		return &object.Boolean{Value: node.Value}


//expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	
//prefix expressions
	case *ast.PrefixExpression:
		right := Eval(node.Right)
		return evalPrefixExpression(node.Operator, right)
	
	return nil
}

//passes evalStatements function and returns Eval(statement)
func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, statement := range stmts {
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
	default:
		return NULL
	}
}