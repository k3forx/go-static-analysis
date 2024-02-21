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

const hello = `
package main

import "fmt"

var x int = 1

func main() {
	var y int = 2
	fmt.Println(x, y)
}
`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", hello, 0)
	if err != nil {
		log.Fatal(err)
	}

	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{InitOrder: []*types.Initializer{}}
	if _, err := conf.Check("cmd/hello", fset, []*ast.File{f}, info); err != nil {
		log.Fatal(err)
	}

	for _, initOrder := range info.InitOrder {
		fmt.Printf("initOrder: %+v\n", initOrder)
		for _, l := range initOrder.Lhs {
			fmt.Printf("Lhs: %s, ", l.Name())
		}
		basiclit := initOrder.Rhs.(*ast.BasicLit)
		fmt.Printf("Rhs: %v\n", basiclit.Value)
	}
}
