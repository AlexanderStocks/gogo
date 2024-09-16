package evaluator

import (
	"fmt"
	"go/ast"
	"go/token"

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
