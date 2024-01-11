package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

const hello = `package main

import "fmt"

func main() {
	fmt.Println("Hello, world")
}`

func main() {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "hello.go", hello, 0)
	if err != nil {
		log.Fatal(err)
	}

	// 型チェックのための設定を行う
	// Importerフィールドはimportの情報を解析するための設定
	conf := types.Config{Importer: importer.Default()}

	// Check関数はtypes.Configに基づいて型チェックを行う
	// 返り値はPackage型になる
	// 最後の引数のtypes.Infoにast.IdentとObjectの紐付けの結果が格納される
	pkg, err := conf.Check("cmd/hello", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Package: %q\n", pkg.Path())
	fmt.Printf("name:    %s\n", pkg.Name())
	fmt.Printf("Imports: %s\n", pkg.Imports())
	fmt.Printf("Scope:   %s\n", pkg.Scope())

	fmt.Println("--------------------------------------")

	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
		Uses: make(map[*ast.Ident]types.Object),
	}

	if _, err := conf.Check("hello", fset, []*ast.File{f}, info); err != nil {
		log.Fatal(err)
	}

	for id, obj := range info.Defs {
		fmt.Printf("%s: %q defines %v\n", fset.Position(id.Pos()), id.Name, obj)
	}

	for id, obj := range info.Uses {
		fmt.Printf("%s: %q uses %v\n", fset.Position(id.Pos()), id.Name, obj)
	}
}
