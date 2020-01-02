package token

// Type represents the type of token for the language specification
type Type string

// Token represents the type of token and the literal string
// representation of that type as defined by the language specification
type Token struct {
	Type    Type
	Literal string
}

const (
	// ILLEGAL represents an illegal character
	ILLEGAL = "ILLEGAL"

	// EOF represents the end of the file
	EOF = "EOF"

	// IDENT represents an identifier
	IDENT = "IDENT" //add, foobar, x, y, ...

	// INT represents an integer
	INT = "INT" //1347686

	// ASSIGN represents the assign operator
	ASSIGN = "="

	// PLUS represents the plus operator
	PLUS = "+"

	// MINUS represents the subtraction operator
	MINUS = "-"

	// ASTERISK represents the multiplaction operator
	ASTERISK = "*"

	// BANG represents the negation operator
	BANG = "!"

	// SLASH represents the forward division operator
	SLASH = "/"

	// GT represents the greater than operator
	GT = ">"

	// LT represents the less than operator
	LT = "<"

	// EQUALS represents the equals operator
	EQUALS = "=="

	// NOTEQUAL represents the not equal operator
	NOTEQUAL = "!="

	// COMMA represents a comma delimiter
	COMMA = ","

	// SEMICOLON represents a semi-colon delimiter
	SEMICOLON = ";"

	// COLON represents a colon delimiter
	COLON = ":"

	// LPAREN represents a left paranthesis delimiter
	LPAREN = "("

	// RPAREN represents a right paranthesis delimiter
	RPAREN = ")"

	// LBRACE represents a left brace delimiter
	LBRACE = "{"

	// RBRACE represents a right brace delimiter
	RBRACE = "{"

	// FUNCTION represents the function keyword
	FUNCTION = "FUNCTION"

	// LET represents the variable assignment keyword
	LET = "LET"

	// TRUE represents the boolean true keyword
	TRUE = "TRUE"

	// FALSE represents the boolean false keyword
	FALSE = "FALSE"

	// IF represents the logical if keyword
	IF = "IF"

	// ELSE represents the logical else keyword
	ELSE = "ELSE"

	// RETURN represents the return keyword
	RETURN = "RETURN"
)

// LookupIdent maps a string identifier to the corresponding token type
func LookupIdent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}
