package evaluator_test

import (
	"go/parser"
	"go/token"
	"testing"

	"github.com/AlexanderStocks/GoGo/internal/evaluator"
	"github.com/AlexanderStocks/GoGo/internal/runtime"
)

func TestEvalExpressions(t *testing.T) {
	tests := []struct {
		input    string
		expected interface{}
	}{
		{"1 + 2", int64(3)},
		{"10 - 5", int64(5)},
		{"2 * 3", int64(6)},
		{"8 / 2", int64(4)},
		{"5 > 2", true},
		{"5 < 2", false},
		{"3 == 3", true},
		{"3 != 4", true},
		{`"hello" + " " + "world"`, "hello world"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := testEvalExpr(tt.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if result != tt.expected {
				t.Errorf("expected %v (%T), got %v (%T)", tt.expected, tt.expected, result, result)
			}
		})
	}
}

func testEvalExpr(input string) (interface{}, error) {
	src := "package main\nfunc main() { var result = " + input + " }"
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
