package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"

	"golang.org/x/tools/go/ssa"
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
	files := []*ast.File{file}

	conf := &types.Config{}
	info := &types.Info{
		Types: make(map[ast.Expr]types.TypeAndValue),
		Defs:  make(map[*ast.Ident]types.Object),
		Uses:  make(map[*ast.Ident]types.Object),
	}

	pkg, err := conf.Check("main", fset, files, info)
	if err != nil {
		log.Fatalf("checking package: %v", err)
	}

	prog := ssa.NewProgram(fset, ssa.BuilderMode(0))
	mainPkg := prog.CreatePackage(pkg, files, info, false)
	mainPkg.Build()

	for _, m := range mainPkg.Members {
		fmt.Printf("member: %+v\n", m)
	}
}
