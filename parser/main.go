package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"log"
)

func main() {
	expr, err := parser.ParseExpr("n+1")
	if err != nil {
		log.Fatal(err)
		return
	}

	binEx, ok := expr.(*ast.BinaryExpr)
	if !ok {
		return
	}

	fmt.Printf("expr: %+v\n", binEx)
}
