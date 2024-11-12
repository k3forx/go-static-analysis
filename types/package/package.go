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

	u := types.Universe
	nilObj := u.Lookup("nil")
	obj, ok := nilObj.(*types.Nil)
	if !ok {
		log.Fatalf("unexpected object type %T", nilObj)
	}
	fmt.Println(obj.Name())
}
