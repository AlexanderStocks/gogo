package evaluator

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/AlexanderStocks/GoGo/internal/runtime"
)

func evalBinaryExpr(expr *ast.BinaryExpr, fset *token.FileSet, env *runtime.Environment) (interface{}, error) {
	left, err := evalExpr(expr.X, fset, env)
	if err != nil {
		return nil, err
	}
	right, err := evalExpr(expr.Y, fset, env)
	if err != nil {
		return nil, err
	}
	switch expr.Op {
	case token.ADD:
		return add(left, right)
	case token.SUB:
		return sub(left, right)
	case token.MUL:
		return mul(left, right)
	case token.QUO:
		return quo(left, right)
	case token.EQL:
		return eql(left, right)
	case token.LSS:
		return lss(left, right)
	case token.GTR:
		return gtr(left, right)
	case token.NEQ:
		return neq(left, right)
	default:
		return nil, fmt.Errorf("unsupported binary operator: %v", expr.Op)
	}
}

// TODO: Use generics to simplify the code

func applyOperation[T any](left, right interface{}, op func(T, T) interface{}) (interface{}, error) {
	l, ok := left.(T)
	if !ok {
		return nil, fmt.Errorf("unsupported left operand type: %T", left)
	}
	r, ok := right.(T)
	if !ok {
		return nil, fmt.Errorf("unsupported right operand type: %T", right)
	}
	return op(l, r), nil
}

func add(left, right interface{}) (interface{}, error) {
	return applyOperation(left, right, func(l, r int64) interface{} {
		return l + r
	})
}

func sub(left, right interface{}) (interface{}, error) {
	return applyOperation(left, right, func(l, r int64) interface{} {
		return l - r
	})
}

func mul(left, right interface{}) (interface{}, error) {
	return applyOperation(left, right, func(l, r int64) interface{} {
		return l * r
	})
}

func quo(left, right interface{}) (interface{}, error) {
	return applyOperation(left, right, func(l, r int64) interface{} {
		return l / r
	})
}

func eql(left, right interface{}) (interface{}, error) {
	return applyOperation(left, right, func(l, r int64) interface{} {
		return l == r
	})
}

func lss(left, right interface{}) (interface{}, error) {
	return applyOperation(left, right, func(l, r int64) interface{} {
		return l < r
	})
}

func gtr(left, right interface{}) (interface{}, error) {
	return applyOperation(left, right, func(l, r int64) interface{} {
		return l > r
	})
}

func neq(left, right interface{}) (interface{}, error) {
	return applyOperation(left, right, func(l, r int64) interface{} {
		return l != r
	})
}
