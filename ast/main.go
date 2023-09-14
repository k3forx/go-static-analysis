package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	const src = `package main
func main() {
;
	goto HOGE
	HOGE:
}
`

	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "", src, parser.Mode(0))
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
			emptyStmt, ok := stmt.(*ast.EmptyStmt)
			if ok {
				fmt.Printf("found EmptyStmt at %v\n", fset.Position(emptyStmt.Pos()))
				ast.Print(fset, emptyStmt)
			}

			labeledStmt, ok := stmt.(*ast.LabeledStmt)
			if ok {
				fmt.Printf("found LabeledStmt at %v\n", fset.Position(labeledStmt.Pos()))
				ast.Print(fset, labeledStmt)
			}
		}
	}
}
