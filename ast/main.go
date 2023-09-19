package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("ast_and_token_pos.go")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer file.Close()

	b, _ := io.ReadAll(file)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file.Name(), b, parser.Mode(0))
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, d := range f.Decls {
		fun, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}

		funcBody := fun.Body
		for _, stmt := range funcBody.List {
			exprStmt, ok := stmt.(*ast.ExprStmt)
			if !ok {
				continue
			}

			sumVar := exprStmt.X.(*ast.BinaryExpr).X.(*ast.Ident)
			fmt.Printf("Pos(): %+v\n", fset.Position(sumVar.Pos()))
			fmt.Printf("End(): %+v\n", fset.Position(sumVar.End()))
			sumVar.
		}
	}
}
