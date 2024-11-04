package info_uses_defs

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"reflect"
)

const hello = `
package main

var a int
const b = 1
type c struct {}

func d() {}
`

func Do() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", hello, parser.Mode(0))
	if err != nil {
		log.Fatalf("parsing file: %v", err)
	}

	conf := types.Config{}
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
	}

	_, err = conf.Check("hello.go", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatalf("checking types: %v", err)
	}

	for ident, o := range info.Defs {
		fmt.Printf("name: %s, types: %v\n", ident.Name, reflect.TypeOf(o))
	}
}
