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

	binaryExpr, ok := (expr).(*ast.BinaryExpr)
	if !ok {
		return
	}

	fmt.Printf("X: %+v\n", binaryExpr.X)
	fmt.Printf("Op: %+v\n", binaryExpr.Op)
	fmt.Printf("Y: %+v\n", binaryExpr.Y)
}
