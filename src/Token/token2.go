//token/token.go

package token

type TokenType string //

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//identifiers and literals
	IDENT = "IDENT" //add, foobar, x, y, ...
	INT   = "INT"   //1347686

	//operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	BANG     = "!"
	SLASH    = "/"
	GT       = ">"
	LT       = "<"
	EQUALS   = "=="
	NOTEQUAL = "!="

	//delimiters
	COMMA     = ","
	SEMICOLON = ";"
	COLON     = ":"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "{"

	//keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

type Token struct {
	Type    TokenType
	Literal string
}
