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
	pkg, err := conf.Check("cmd/hello", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Package: %q\n", pkg.Path())
	fmt.Printf("name:    %s\n", pkg.Name())
	fmt.Printf("Imports: %s\n", pkg.Imports())
	fmt.Printf("Scope:   %s\n", pkg.Scope())
	fmt.Println(pkg.Scope().Lookup("main").Type())
	fmt.Println(pkg.Imports()[0])
}
