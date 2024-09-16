package runtime

import "fmt"

type ObjectType string

const (
	IntegerObj = "INTEGER"
	BooleanObj = "BOOLEAN"
	StringObj  = "STRING"
	// Define other object types as needed.
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() ObjectType { return IntegerObj }
func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BooleanObj }
func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }

type String struct {
	Value string
}

func (s *String) Type() ObjectType { return StringObj }
func (s *String) Inspect() string  { return s.Value }

// Define other object types like Function, ReturnValue, Error, etc.
