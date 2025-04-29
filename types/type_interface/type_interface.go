package type_interface

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"slices"
)

const pkgMain = `package main

// types.Basic
var a int

// types.Pointer
var b *float64

// types.Array
var c [0]string

// types.Slice
var d []int

// types.Map
var e map[string]int

// types.Chan
var f chan int

// types.Struct
var g struct{}

// types.Signature
func h()

// types.Alias
type i = int

// types.Named
type j struct{}

// types.Interface
var k any

// TypeParam
func someFunc[l any]() {}

// types.Tuple
func someFunc2(str string) (int, bool)

// types.Union
type someInterface interface {
	int | string
}
`

func Do() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "main.go", pkgMain, parser.Mode(0))
	if err != nil {
		log.Fatalf("parsing file: %v", err)
	}

	conf := types.Config{}
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
	}
	_, err = conf.Check("main.go", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatalf("checking types: %v", err)
	}

	for ident, obj := range info.Defs {
		// `package main` はobjectがnilになるのでスキップ
		if obj == nil {
			continue
		}

		// if ident.Name == "m" {
		// 	tp, ok := obj.Type().Underlying().(*types.Interface)
		// 	if !ok {
		// 		continue
		// 	}
		// 	for i := range tp.NumEmbeddeds() {
		// 		fmt.Printf("embedded: %s\n", reflect.TypeOf(tp.EmbeddedType(i)))
		// 	}
		// }

		// if ident.Name == "h" {
		// 	tp, ok := obj.Type().(*types.Signature)
		// 	if !ok {
		// 		continue
		// 	}
		// 	fmt.Printf("type of result of func 'h': %+v\n", reflect.TypeOf(tp.Results()))
		// }

		if !slices.Contains([]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n"}, ident.Name) {
			continue
		}

		switch tp := obj.Type().(type) {
		case *types.Basic:
			fmt.Printf("type of '%s' is types.Basic (%s)\n", ident.Name, tp.String())
		case *types.Pointer:
			fmt.Printf("type of '%s' is types.Pointer (%s)\n", ident.Name, tp.String())
		case *types.Array:
			fmt.Printf("type of '%s' is types.Array (%s)\n", ident.Name, tp.String())
		case *types.Slice:
			fmt.Printf("type of '%s' is types.Slice (%s)\n", ident.Name, tp.String())
		case *types.Map:
			fmt.Printf("type of '%s' is types.Map (%s)\n", ident.Name, tp.String())
		case *types.Chan:
			fmt.Printf("type of '%s' is types.Chan (%s)\n", ident.Name, tp.String())
		case *types.Struct:
			fmt.Printf("type of '%s' is types.Struct (%s)\n", ident.Name, tp.String())
		case *types.Tuple:
			fmt.Printf("type of '%s' is types.Tuple (%s)\n", ident.Name, tp.String())
		case *types.Signature:
			fmt.Printf("type of '%s' is types.Signature (%s)\n", ident.Name, tp.String())
		case *types.Alias:
			fmt.Printf("type of '%s' is types.Alias (%s)\n", ident.Name, tp.String())
		case *types.Named:
			fmt.Printf("type of '%s' is types.Named (%s)\n", ident.Name, tp.String())
		case *types.Interface:
			fmt.Printf("type of '%s' is types.Interface (%s)\n", ident.Name, tp.String())
		case *types.Union:
			fmt.Printf("type of '%s' is types.Union (%s)\n", ident.Name, tp.String())
		case *types.TypeParam:
			fmt.Printf("type of '%s' is types.TypeParam (%s)\n", ident.Name, tp.String())
		}
	}
}
