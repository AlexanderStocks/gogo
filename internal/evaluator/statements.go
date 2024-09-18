package evaluator

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/AlexanderStocks/GoGo/internal/runtime"
)

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
	case *ast.AssignStmt:
		return evalAssignStmt(s, fset, env)
	case *ast.DeclStmt:
		return evalDeclStmt(s, fset, env)
	case *ast.IfStmt:
		return evalIfStmt(s, fset, env)
	case *ast.ForStmt:
		return evalForStmt(s, fset, env)
	// case *ast.SelectStmt:
	// 	return evalSelectStmt(s, fset, env)
	// case *ast.SwitchStmt:
	// 	return evalSwitchStmt(s, fset, env)
	// case *ast.BranchStmt:
	// 	return evalBranchStmt(s, fset, env)
	// case *ast.LabeledStmt:
	// 	return evalLabeledStmt(s, fset, env)
	// case *ast.GoStmt:
	// 	return evalGoStmt(s, fset, env)
	// case *ast.DeferStmt:
	// 	return evalDeferStmt(s, fset, env)
	// case *ast.ReturnStmt:
	// 	return evalReturnStmt(s, fset, env)
	// case *ast.IncDecStmt:
	// 	return evalIncDecStmt(s, fset, env)
	// case *ast.SendStmt:
	// 	return evalSendStmt(s, fset, env)
	// case *ast.EmptyStmt:
	// 	return evalEmptyStmt(s, fset, env)
	default:
		return fmt.Errorf("unsupported statement type: %T", s)
	}
}
