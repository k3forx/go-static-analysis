package object

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"reflect"
)

const pkgMain = `package main

import _ "fmt"

const a = "hello"
var b *int
type c struct {}

func main() {
LOOP:
	for {
		break LOOP
	}
}
`

func Do() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", pkgMain, parser.Mode(0))
	if err != nil {
		log.Fatalf("parsing file: %v", err)
	}

	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		// Defs: make(map[*ast.Ident]types.Object),
		Uses: make(map[*ast.Ident]types.Object),
	}
	_, err = conf.Check("main.go", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatalf("checking types: %v", err)
	}

	for ident, obj := range info.Defs {
		fmt.Printf("%s: %+v\n", fset.Position(ident.Pos()), reflect.TypeOf(obj))
	}

	println()

	for ident, obj := range info.Uses {
		fmt.Println(ident.Name)
		fmt.Printf("%s: %+v\n", fset.Position(ident.Pos()), reflect.TypeOf(obj))
	}
}
