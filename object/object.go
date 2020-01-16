package object


import "fmt"

//defined ObjectType type
type ObjectType string

type Object interface {
	Type() ObjectType
	Inspect() string
}

//defined constants of INTEGER_OBJ, BOOLEAN_OBJ, and NULL_OBJ
const (
	INTEGER_OBJ = "INTEGER"
)

const (
	BOOLEAN_OBJ = "BOOLEAN"
)

const (
	NULL_OBJ = "NULL"
)

//defined null, boolean, and integer types
type Null struct {}

type Boolean struct {
	Value bool
}

type Integer struct {
	Value int64
}

//passed inspect function for integers
func (i *Integer) Inspect() string { return fmt.Sprintf("%d, i.Value") }

//passed type function for integers
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

//passed inspect function for booleans
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t, b.Value") }

//passed type function for booleans
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

//passed inspect function for null
func (n *Null) Inspect() string { return "null"}

//passed type function for null
func (n *Null) Type() ObjectType { return NULL_OBJ }



