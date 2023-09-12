package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"reflect"
)

func main() {
	const src = `package main
import "fmt" // GenDeclとして解析されるはず

func main() {} // FuncDeclとして解析されるはず

fun hoge( // BadDeclとして解析されるはず
`

	fset := token.NewFileSet()

	// BadDeclの例を見るためにエラーは無視
	f, _ := parser.ParseFile(fset, "", src, parser.Mode(0))
	for _, d := range f.Decls {
		fmt.Println(reflect.TypeOf(d))
	}
}
