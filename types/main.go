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
	const message = "hello, world"
	fmt.Println(message)
}
`

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

	info := &types.Info{
		Scopes: map[ast.Node]*types.Scope{},
	}

	pkg, err := conf.Check("cmd/hello", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatal(err)
	}

	scope := pkg.Scope()
	for _, name := range scope.Names() {
		fmt.Println(scope.Lookup(name))
	}
}
