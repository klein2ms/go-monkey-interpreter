import "fmt"

//type method that return its ObjectType
type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
)

type Interger struct {
	Value int64
}

func (i *Integer) Inspect() string { return fmt.Sprintf("%d, i.Value") }

func (i *Integer) Type() ObjectType { return INTEGER_OBJ }
