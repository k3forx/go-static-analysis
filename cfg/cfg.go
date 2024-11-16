package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"

	"golang.org/x/tools/go/cfg"
)

const pkgMain = `package main

func main() {
	var a int = 10
	if a > 5 {
		a = 5
	}
}
`

func main() {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "main.go", pkgMain, parser.Mode(0))
	if err != nil {
		log.Fatalf("parsing file: %v", err)
	}

	ast.Inspect(file, func(n ast.Node) bool {
		funcDecl, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		if funcDecl.Name.Name != "main" {
			return true
		}

		// ast.Print(fset, funcDecl)

		graph := cfg.New(funcDecl.Body, func(ce *ast.CallExpr) bool {
			return false
		})
		fmt.Println(graph.Format(fset))
		fmt.Println("all blocks")
		for _, block := range graph.Blocks {
			fmt.Printf("block: %+v\n", block)
		}
		println()

		block := graph.Blocks[0]

		fmt.Printf("block: %+v\n", block)
		for _, node := range block.Nodes {
			fmt.Printf("node: %T\n", node)
		}

		for _, succ := range block.Succs {
			fmt.Printf("successor: %+v\n", succ)
		}

		fmt.Printf("return: %+v\n", block.Return())

		ast.Print(fset, block.Stmt)
		// fmt.Printf("stmt: %T\n", block.Stmt)
		// stmt := block.Stmt.(*ast.BlockStmt)
		// for _, stmt := range stmt.List {
		// 	fmt.Printf("stmt: %T\n", stmt)
		// }

		return true
	})
}
