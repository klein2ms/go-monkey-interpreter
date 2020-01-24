package evaluator

import (
	"github.com/klein2ms/go-monkey-interpreter/object"
	"github.com/klein2ms/go-monkey-interpreter/ast"
)

//Eval function for object.Object
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

//statements
	case *ast.Program:
		return evalStatements(node.Statements)
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

//expressions
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}
	
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