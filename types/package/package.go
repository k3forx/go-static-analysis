package pkg

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

const pkgStr = `package main
var a string

func main() {
	var b int
	b = 123
	println(b)
}
`

func Do() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", pkgStr, parser.Mode(0))
	if err != nil {
		log.Fatalf("parsing file: %v", err)
	}

	conf := types.Config{}
	pkgInfo, err := conf.Check("main.go", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatalf("checking types: %v", err)
	}

	fmt.Printf("Name: %s\n", pkgInfo.Name())
	fmt.Printf("Scope: %+v\n", pkgInfo.Scope().Names())
	for i := range pkgInfo.Scope().NumChildren() {
		sc := pkgInfo.Scope().Child(i)
		fmt.Printf("Child: %+v\n", sc.Child(0))
	}
}
