package ast

import (
	"github.com/klein2ms/go-monkey-interpreter/token"
)

// Node represents a node in the AST
type Node interface {
	TokenLiteral() string
}

// Statement represents a statement in the AST
type Statement interface {
	Node
	statementNode()
}

// Expression represents an expression in the AST
type Expression interface {
	Node
	expressionNode()
}

// Program represents a program composed of zero or many Statements
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the TokenLiteral root statement in Program
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// LetStatement represents a statement for the Let keyword in the AST
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {

}

// TokenLiteral returns the TokenLiteral for a LetStatement
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier represents an identifier in the AST
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {

}

// TokenLiteral returns the TokenLiteral for an Idenfitier
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// ReturnStatement represents a statement for the Return keyword in the AST
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {

}

// TokenLiteral returns the TokenLiteral for a ReturnStatement
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}
