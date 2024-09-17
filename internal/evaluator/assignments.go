package evaluator

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/AlexanderStocks/GoGo/internal/runtime"
)

func evalAssignStmt(stmt *ast.AssignStmt, fset *token.FileSet, env *runtime.Environment) error {
	if len(stmt.Lhs) != len(stmt.Rhs) {
		return fmt.Errorf("assignment count mismatch: %d = %d", len(stmt.Lhs), len(stmt.Rhs))
	}

	for i, lhs := range stmt.Lhs {
		ident, ok := lhs.(*ast.Ident)
		if !ok {
			return fmt.Errorf("unsupported assignment target: %T", lhs)
		}

		val, err := evalExpr(stmt.Rhs[i], fset, env)
		if err != nil {
			return err
		}

		env.Set(ident.Name, val)
	}

	return nil
}
