package evaluator

import (
	"go/ast"
	"go/token"

	"github.com/AlexanderStocks/GoGo/internal/runtime"
)

func evalForStmt(s *ast.ForStmt, fset *token.FileSet, env *runtime.Environment) error {
	if s.Init != nil {
		err := evalStmt(s.Init, fset, env)
		if err != nil {
			return err
		}
	}

	for {
		cond, err := evalExpr(s.Cond, fset, env)
		if err != nil {
			return err
		}

		if !cond.(bool) {
			break
		}

		err = evalBlockStmt(s.Body, fset, env)
		if err != nil {
			return err
		}

		if s.Post != nil {
			err = evalStmt(s.Post, fset, env)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
