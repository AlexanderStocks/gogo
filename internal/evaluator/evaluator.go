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
	// case *ast.DeclStmt:
	// 	return evalDeclStmt(s, fset, env)
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
	// case *ast.BinaryExpr:
	//     return evalBinaryExpr(e, fset, env)
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
