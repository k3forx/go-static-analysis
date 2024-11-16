package info_selections

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

const pkgMain = `package main
type T struct{Field int}
func (T) Method() {}
var v T

                 // Kind            Type
var _ = v.Field  // FieldVal        int
var _ = v.Method // MethodVal       func()
var _ = T.Method // MethodExpr      func(T)

func main() {
	v.Method()  // FieldVal        int
}
`

func Do() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.go", pkgMain, parser.Mode(0))
	if err != nil {
		log.Fatalf("parsing file: %v", err)
	}

	conf := types.Config{}
	info := &types.Info{
		Selections: make(map[*ast.SelectorExpr]*types.Selection),
	}

	_, err = conf.Check("main.go", fset, []*ast.File{file}, info)
	if err != nil {
		log.Fatalf("checking types: %v", err)
	}

	for selectorExpr, selection := range info.Selections {
		fmt.Printf("selector: %s, selection: %+v\n", selectorExpr.Sel.Name, selection.String())
	}
}
