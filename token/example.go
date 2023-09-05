package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func example() {
	fset := token.NewFileSet()

	const src = `package main

import "fmt"

import "go/token"

//line :1:5
type p = token.Pos

const bad = token.NoPos

//line fake.go:42:11
func ok(pos p) bool {
		return pos != bad
}

/*line :7:9*/func main() {
	fmt.Println(ok(bad) == bad.IsValid())
}
`

	f, err := parser.ParseFile(fset, "main.go", src, 0)
	if err != nil {
		fmt.Printf("err: %+v\n", err)
		return
	}

	// fのそれぞれの宣言の位置と種類を出力する
	for _, decl := range f.Decls {
		pos := decl.Pos()
		relPosition := fset.Position(pos)
		absPosition := fset.PositionFor(pos, false)

		kind := "func"
		if gen, ok := decl.(*ast.GenDecl); ok {
			kind = gen.Tok.String()
		}

		// If the relative and absolute positions differ, show both.
		fmtPosition := relPosition.String()
		if relPosition != absPosition {
			fmtPosition += "[" + absPosition.String() + "]"
		}

		fmt.Printf("%s: %s\n", fmtPosition, kind)
	}
}
