package token

type TokenType string //
  
const (
	ILLEGAL = "ILLEGAL"
	EOF   "EOF"
 
	//identifiers and literals
	IDENT = "IDENT" //add, foobar, x, y, ...
	INT   = "INT"   //1347686
   
	//oper  tors
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	AS      RISK  = "*"
	BANG      = "!"
	SLASH     = "/"
	GT        = ">"
	LT        = "<"
	EQUALS    = "=="
	NOTEQUAL = "!="

	//del   miters
	COMMA     = ","
	SEMICOLON  = ";"
	COLON     = ":"
	LPAREN     = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE   = "{"
   
	//kewords
	FUNCTION  = "FUNCTION"
	LET      = "LET"
	TRUE       = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)    

	Type	TokenType

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"let": LET,
}

func LookupIdent( ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return token
	}
	return IDENT
}