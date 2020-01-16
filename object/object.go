package object

import "fmt"

// Type represents the type of the object representation in the language specification
type Type string

const (
	// INTEGER represents an integer type
	INTEGER = "INTEGER"
	// BOOLEAN represents a boolean type
	BOOLEAN = "BOOLEAN"
	// NULL represents a null or nil type
	NULL = "NULL"
)

// Object represents the interface by which each object in the language specification must abide
type Object interface {
	Type() Type
	Inspect() string
}

// Integer represents an integer in the language specification
type Integer struct {
	Value int64
}

// Inspect returns the value of the integer.
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type returns the Integer Type
func (i *Integer) Type() Type { return INTEGER }

// Boolean represents a boolean in the language specification
type Boolean struct {
	Value bool
}

// Inspect returns the value of the boolean.
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Type returns the Boolean Type
func (b *Boolean) Type() Type { return BOOLEAN }

// Null represents a null in the language specification
type Null struct{}

// Inspect returns 'null'.
func (n *Null) Inspect() string { return "null" }

// Type returns the Null Type
func (n *Null) Type() Type { return NULL }
