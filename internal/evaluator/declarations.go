package evaluator

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/AlexanderStocks/GoGo/internal/runtime"
)

func evalDeclStmt(declStmt *ast.DeclStmt, fset *token.FileSet, env *runtime.Environment) error {
	genDecl, ok := declStmt.Decl.(*ast.GenDecl)
	if !ok {
		return fmt.Errorf("unsupported declaration: %T", declStmt.Decl)
	}
	for _, spec := range genDecl.Specs {
		err := evalSpec(spec, fset, env)
		if err != nil {
			return err
		}
	}
	return nil
}

func evalSpec(spec ast.Spec, fset *token.FileSet, env *runtime.Environment) error {
	switch s := spec.(type) {
	case *ast.ValueSpec:
		for i, name := range s.Names {
			var value interface{}
			if i < len(s.Values) {
				val, err := evalExpr(s.Values[i], fset, env)
				if err != nil {
					return err
				}
				value = val
			} else {
				value = getZeroValue(s.Type)
			}
			env.Set(name.Name, value)
		}
		return nil
	default:
		return fmt.Errorf("unsupported spec type: %T", s)
	}
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
