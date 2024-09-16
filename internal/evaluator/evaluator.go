package evaluator

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"

	"github.com/AlexanderStocks/GoGo/internal/runtime"
)

func Eval(node ast.Node, fset *token.FileSet, env *runtime.Environment) error {
	switch n := node.(type) {
	case *ast.File:
		for _, decl := range n.Decls {
			err := Eval(decl, fset, env)
			if err != nil {
				return err
			}
		}
	case *ast.FuncDecl:
		if n.Name.Name == "main" {
			return Eval(n.Body, fset, env)
		}
		// Store function declarations in the environment if needed.
	case *ast.BlockStmt:
		return evalBlockStmt(n, fset, env)
	case ast.Stmt:
		return evalStmt(n, fset, env)
	case ast.Expr:
		_, err := evalExpr(n, fset, env)
		return err
	default:
		return fmt.Errorf("unsupported node type: %T", n)
	}
	return nil
}

func evalBlockStmt(block *ast.BlockStmt, fset *token.FileSet, env *runtime.Environment) error {
	for _, stmt := range block.List {
		err := evalStmt(stmt, fset, env)
		if err != nil {
			return err
		}
	}
	return nil
}

func evalStmt(stmt ast.Stmt, fset *token.FileSet, env *runtime.Environment) error {
	switch s := stmt.(type) {
	case *ast.ExprStmt:
		res, err := evalExpr(s.X, fset, env)
		if err == nil {
			fmt.Println(res)
		}
		return err
	// case *ast.AssignStmt:
	// 	return evalAssignStmt(s, fset, env)
	case *ast.DeclStmt:
		return evalDeclStmt(s, fset, env)
	// case *ast.IfStmt:
	// 	return evalIfStmt(s, fset, env)
	// case *ast.ForStmt:
	// 	return evalForStmt(s, fset, env)
	default:
		return fmt.Errorf("unsupported statement type: %T", s)
	}
}

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

// Implement evaluation functions for expressions and statements...

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

func evalDeclStmt(declStmt *ast.DeclStmt, fset *token.FileSet, env *runtime.Environment) error {
	genDecl, ok := declStmt.Decl.(*ast.GenDecl)
	if !ok {
		return fmt.Errorf("unsupported declaration: %T", declStmt.Decl)
	}
	for _, spec := range genDecl.Specs {
		valueSpec, ok := spec.(*ast.ValueSpec)
		if !ok {
			return fmt.Errorf("unsupported spec: %T", spec)
		}
		for i, name := range valueSpec.Names {
			var value interface{}
			if i < len(valueSpec.Values) {
				val, err := evalExpr(valueSpec.Values[i], fset, env)
				if err != nil {
					return err
				}
				value = val
			} else {
				value = getZeroValue(valueSpec.Type)
			}
			env.Set(name.Name, value)
		}
	}
	return nil
}

func getZeroValue(expr ast.Expr) interface{} {
	switch expr.(type) {
	case *ast.Ident:
		// Handle built-in types
		return 0
	default:
		return nil
	}
}

func evalSpec(spec ast.Spec, fset *token.FileSet, env *runtime.Environment) (interface{}, error) {
	switch s := spec.(type) {
	case *ast.ValueSpec:
		if len(s.Values) == 0 {
			return nil, nil
		}
		return evalExpr(s.Values[0], fset, env)
	default:
		return nil, fmt.Errorf("unsupported spec type: %T", s)
	}
}

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
	// case token.SUB:
	// 	return sub(left, right)
	// case token.MUL:
	// 	return mul(left, right)
	// case token.QUO:
	// 	return quo(left, right)
	// case token.EQL:
	// 	return eql(left, right)
	// case token.LSS:
	// 	return lss(left, right)
	// case token.GTR:
	// 	return gtr(left, right)
	default:
		return nil, fmt.Errorf("unsupported binary operator: %v", expr.Op)
	}
}

func add(left, right interface{}) (interface{}, error) {
	switch l := left.(type) {
	case int64:
		switch r := right.(type) {
		case int64:
			return l + r, nil
		default:
			return nil, fmt.Errorf("unsupported right operand type: %T", right)
		}
	default:
		return nil, fmt.Errorf("unsupported left operand type: %T", left)
	}
}
