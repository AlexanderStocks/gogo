package evaluator

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"

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
	return applyOperation(left, right, expr.Op)
}

func applyOperation(left, right interface{}, op token.Token) (interface{}, error) {
	lv := reflect.ValueOf(left)
	rv := reflect.ValueOf(right)

	if lv.Kind() != rv.Kind() {
		return nil, fmt.Errorf("type mismatch: %s vs %s", lv.Kind(), rv.Kind())
	}

	switch lv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return applyIntOp(lv.Int(), rv.Int(), op)
	case reflect.Float32, reflect.Float64:
		return applyFloatOp(lv.Float(), rv.Float(), op)
	case reflect.String:
		return applyStringOp(lv.String(), rv.String(), op)
	default:
		return nil, fmt.Errorf("unsupported type: %s", lv.Kind())
	}
}

func applyIntOp(left, right int64, op token.Token) (interface{}, error) {
	switch op {
	case token.ADD:
		return left + right, nil
	case token.SUB:
		return left - right, nil
	case token.MUL:
		return left * right, nil
	case token.QUO:
		if right == 0 {
			return nil, fmt.Errorf("division by zero")
		}
		return left / right, nil
	case token.EQL:
		return left == right, nil
	case token.NEQ:
		return left != right, nil
	case token.LSS:
		return left < right, nil
	case token.GTR:
		return left > right, nil
	default:
		return nil, fmt.Errorf("unsupported operator for ints: %v", op)
	}
}

func applyFloatOp(left, right float64, op token.Token) (interface{}, error) {
	switch op {
	case token.ADD:
		return left + right, nil
	case token.SUB:
		return left - right, nil
	case token.MUL:
		return left * right, nil
	case token.QUO:
		if right == 0.0 {
			return nil, fmt.Errorf("division by zero")
		}
		return left / right, nil
	case token.EQL:
		return left == right, nil
	case token.NEQ:
		return left != right, nil
	case token.LSS:
		return left < right, nil
	case token.GTR:
		return left > right, nil
	default:
		return nil, fmt.Errorf("unsupported operator for floats: %v", op)
	}
}

func applyStringOp(left, right string, op token.Token) (interface{}, error) {
	switch op {
	case token.ADD:
		return left + right, nil
	case token.EQL:
		return left == right, nil
	case token.NEQ:
		return left != right, nil
	case token.LSS:
		return left < right, nil
	case token.GTR:
		return left > right, nil
	default:
		return nil, fmt.Errorf("unsupported operator for strings: %v", op)
	}
}

func evalUnaryExpr(expr *ast.UnaryExpr, fset *token.FileSet, env *runtime.Environment) (interface{}, error) {
	x, err := evalExpr(expr.X, fset, env)
	if err != nil {
		return nil, err
	}
	return applyUnaryOp(x, expr.Op)
}

func applyUnaryOp(x interface{}, op token.Token) (interface{}, error) {
	v := reflect.ValueOf(x)

	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return applyIntUnaryOp(v.Int(), op)
	case reflect.Float32, reflect.Float64:
		return applyFloatUnaryOp(v.Float(), op)
	default:
		return nil, fmt.Errorf("unsupported type: %s", v.Kind())
	}
}

func applyIntUnaryOp(x int64, op token.Token) (interface{}, error) {
	switch op {
	case token.SUB:
		return -x, nil
	default:
		return nil, fmt.Errorf("unsupported operator for ints: %v", op)
	}
}

func applyFloatUnaryOp(x float64, op token.Token) (interface{}, error) {
	switch op {
	case token.SUB:
		return -x, nil
	default:
		return nil, fmt.Errorf("unsupported operator for floats: %v", op)
	}
}
