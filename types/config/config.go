package config

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

const pkgStr = `package main

import (
	"fmt"
	_ "context"
	. "errors"
)

func main() {
	fmt.Println("Hello, World!")
	New("test")
}
`

func Do() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", pkgStr, parser.Mode(0))
	if err != nil {
		log.Fatalf("parsing file: %v", err)
	}

	conf := types.Config{
		Importer: importer.Default(),
	}
	pkgInfo, err := conf.Check("main.go", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatalf("checking types: %v", err)
	}

	fmt.Printf("Name: %s\n", pkgInfo.Name())
}
