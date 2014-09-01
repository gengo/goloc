// package goloc counts the number of statements in a Go source file
package goloc

import (
	"go/ast"
	"go/parser"
	"go/token"
)

// CountStatements counts the number of statements in a Go source file, given a string
// representation of the file. It returns an error if the string is not a properly-formatted
// Go source file.
func CountStatements(src string) (int, error) {
	// Create the AST by parsing src.
	fset := token.NewFileSet() // positions are relative to fset
	f, err := parser.ParseFile(fset, "", src, 0)
	if err != nil {
		return 0, err
	}

	stmts := 0
	prevLine := 0

	// Inspect the AST and print all identifiers and literals.
	ast.Inspect(f, func(n ast.Node) bool {
		switch n.(type) {
		case *ast.Comment, *ast.CommentGroup, *ast.GenDecl, *ast.Package:
			return false
		case *ast.FuncDecl, *ast.FuncType, *ast.FuncLit:
			return true
		case *ast.AssignStmt, *ast.GoStmt, *ast.IfStmt, *ast.ForStmt,
			*ast.DeclStmt, *ast.ExprStmt, *ast.SendStmt,
			*ast.SelectStmt, *ast.RangeStmt,
			*ast.SwitchStmt, *ast.IncDecStmt, *ast.LabeledStmt,
			*ast.ReturnStmt, *ast.TypeSwitchStmt:
			if n != nil {
				l := fset.Position(n.Pos()).Line
				if prevLine != l {
					prevLine = l
					stmts += 1
				}
			}
			return true
		}

		return true
	})

	return stmts, nil
}
