package interpreter

import (
	"go/parser"
	"go/token"

	"github.com/AlexanderStocks/GoGo/internal/evaluator"
	"github.com/AlexanderStocks/GoGo/internal/runtime"
)

func RunFile(filename string) error {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, filename, nil, 0)
	if err != nil {
		return err
	}

	env := runtime.NewEnvironment(nil)

	err = evaluator.Eval(file, fset, env)
	if err != nil {
		return err
	}

	return nil
}
