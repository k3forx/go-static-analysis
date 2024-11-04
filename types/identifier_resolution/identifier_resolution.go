package identifier_resolution

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

const hello = `
package main

const (
	a = 1
	b = 2
	sum = a + b
)

func main() {
	var x int
	println(x)

	y := "hello"
	println(y)

	println(sum)
}
`

func Do() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", hello, parser.Mode(0))
	if err != nil {
		log.Fatalf("parsing file: %w", err)
	}

	conf := types.Config{}
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
		Uses: make(map[*ast.Ident]types.Object),
	}

	_, err = conf.Check("hello.go", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatalf("checking types: %w", err)
	}

	// printlnで使用されているxはどこで定義されているか？ -> 識別子の解決
	// yの型は何か？ -> 型推論
	// sumが定義できるか？ -> 定数式の評価
	for _, obj := range info.Uses {
		for defIdent, o := range info.Defs {
			if o == obj {
				switch defIdent.Name {
				case "x":
					fmt.Println("xの型は", o.Type().String(), o.String())
				case "y":
					fmt.Println("yの型は", o.Type().String())
				case "sum":
					fmt.Println("sumの型は", o.Type().String())
				}
			}
		}
	}
}
