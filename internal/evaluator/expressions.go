package evaluator

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"

	"github.com/AlexanderStocks/GoGo/internal/runtime"
)

func evalExpr(expr ast.Expr, fset *token.FileSet, env *runtime.Environment) (interface{}, error) {
	switch e := expr.(type) {
	case *ast.BasicLit:
		return evalBasicLit(e)
	case *ast.Ident:
		if e.Name == "println" {
			return func(args ...interface{}) (interface{}, error) {
				fmt.Println(args...)
				return nil, nil
			}, nil
		}
		return env.Get(e.Name)
	case *ast.BinaryExpr:
		return evalBinaryExpr(e, fset, env)
	case *ast.CallExpr:
		return evalCallExpr(e, fset, env)
	// case *ast.UnaryExpr:
	//     return evalUnaryExpr(e, fset, env)
	default:
		return nil, fmt.Errorf("unsupported expression type: %T", e)
	}
}

func evalBasicLit(lit *ast.BasicLit) (interface{}, error) {
	switch lit.Kind {
	case token.INT:
		// Parse integer literals
		value, err := strconv.ParseInt(lit.Value, 0, 64)
		if err != nil {
			return nil, err
		}
		return value, nil
	case token.STRING:
		// Unquote string literals
		value, err := strconv.Unquote(lit.Value)
		if err != nil {
			return nil, err
		}
		return value, nil
	case token.FLOAT:
		// Parse float literals
		value, err := strconv.ParseFloat(lit.Value, 64)
		if err != nil {
			return nil, err
		}
		return value, nil
	default:
		return nil, fmt.Errorf("unsupported literal kind: %v", lit.Kind)
	}
}

func evalCallExpr(call *ast.CallExpr, fset *token.FileSet, env *runtime.Environment) (interface{}, error) {
	fun, err := evalExpr(call.Fun, fset, env)
	if err != nil {
		return nil, err
	}
	args := []interface{}{}
	for _, arg := range call.Args {
		val, err := evalExpr(arg, fset, env)
		if err != nil {
			return nil, err
		}
		args = append(args, val)
	}
	switch fn := fun.(type) {
	case func(args ...interface{}) (interface{}, error):
		return fn(args...)
	default:
		return nil, fmt.Errorf("attempt to call non-function: %T", fun)
	}
}
