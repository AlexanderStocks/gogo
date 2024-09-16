package runtime

import "fmt"

type Environment struct {
	store map[string]interface{}
	outer *Environment
}

func NewEnvironment(outer *Environment) *Environment {
	return &Environment{
		store: make(map[string]interface{}),
		outer: outer,
	}
}

func (e *Environment) Get(name string) (interface{}, error) {
	val, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.Get(name)
	}
	if !ok {
		return nil, fmt.Errorf("undefined variable: %s", name)
	}
	return val, nil
}

func (e *Environment) Set(name string, val interface{}) {
	e.store[name] = val
}
