package ast

import (
	"github.com/klein2ms/go-monkey-interpreter/token"
	"bytes"
)

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type Node interface {
	TokenLiteral() string
	String() string
}

// The first token of the expression
type ExpressionStatement struct {
	Token		token.Token
	Expression  Expression
}

// the token.LET token
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// the token.IDENT token
type Identifier struct {
	Token token.Token
	Value string
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (es *ExpressionStatement) statementNode()		{}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
func (ls *LetStatement) statementNode()       {}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func ReturnStatement struct {
	Token			token.Token  //  the return token
	ReturnValue 	Expression
}

func (rs *ReturnStatement) statementNode()		{}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.string())
	}

	return out.String()
}

func (i *Identifier) expressionNode()      {}
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// Defining necessary structure to represent return statements.
type ReturnStatement struct {
	Token       token.Token // the 'return' token
	ReturnValue expression
}

func (rs *ReturnStatement) statementNode()       {}
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
