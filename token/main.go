package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"log"
)

func main() {
	fset := token.NewFileSet()

	_, err := parser.ParseDir(fset, "./tmp", nil, 0)
	if err != nil {
		log.Fatal(err)
		return
	}

	pos := token.Pos(13)
	fmt.Printf("position: %+v\n", fset.Position(pos))
}
