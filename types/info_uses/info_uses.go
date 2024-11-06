package info_uses

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

const hello = `package main

type a struct {
	id int // intが使われている
}

func b () {
	var c a // aが使われている
	c.id = 1 // cが使われている、idが使われている
}
`

func Do() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", hello, parser.Mode(0))
	if err != nil {
		log.Fatalf("parsing file: %v", err)
	}

	conf := types.Config{}
	info := &types.Info{
		Uses: make(map[*ast.Ident]types.Object),
	}

	_, err = conf.Check("hello.go", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatalf("checking types: %v", err)
	}

	for ident, o := range info.Uses {
		fmt.Printf("name: %s, types: %v (%s)\n", ident.Name, o.Type().String(), fset.Position(ident.Pos()))
	}
}
