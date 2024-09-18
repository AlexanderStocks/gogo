package evaluator

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/AlexanderStocks/GoGo/internal/runtime"
)

func evalIfStmt(stmt *ast.IfStmt, fset *token.FileSet, env *runtime.Environment) error {
	if stmt.Init != nil {
		err := evalStmt(stmt.Init, fset, env)
		if err != nil {
			return err
		}
	}

	cond, err := evalExpr(stmt.Cond, fset, env)

	if err != nil {
		return err
	}

	if cond.(bool) {
		return evalBlockStmt(stmt.Body, fset, env)
	} else if stmt.Else != nil {
		switch s := stmt.Else.(type) {
		case *ast.BlockStmt:
			return evalBlockStmt(s, fset, env)
		case *ast.IfStmt:
			return evalIfStmt(s, fset, env)
		default:
			return fmt.Errorf("unsupported else type: %T", s)
		}
	}

	return nil
}
