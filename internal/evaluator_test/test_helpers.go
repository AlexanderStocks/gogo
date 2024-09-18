package evaluatortest

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/AlexanderStocks/GoGo/internal/evaluator"
	"github.com/AlexanderStocks/GoGo/internal/runtime"
)

func testEval(t *testing.T, input string) (interface{}, error) {
	t.Helper()
	src := "package main\nfunc main() {\n" + input + "\n}"
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "test.go", src, 0)
	if err != nil {
		return nil, err
	}

	env := runtime.NewEnvironment(nil)
	err = evaluator.Eval(file, fset, env)
	if err != nil {
		return nil, err
	}

	return env.Get("result")
}
