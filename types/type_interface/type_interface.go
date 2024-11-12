package type_interface

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"reflect"
)

const pkgMain = `package main

// types.Basic
// var a int

// types.Pointer
// var b *float64

// types.Array
// var c [0]string

// types.Slice
// var d []int

// types.Map
// var e map[string]int

// types.Chan
// var f chan int

// types.Struct
// var g struct{}

// types.Tuple
func h(k string) (int, bool) 

// types.Signature
// func i()

// types.Alias
// type j = int

// types.Named
// type k struct{}

// types.Interface
// var l interface{}

// types.Union
// type Union interface {
// 	int | string
// }

// TypeParam
// func hoge[T any]() {}
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
		if ident.Name == "Union" {
			tp, ok := obj.Type().Underlying().(*types.Interface)
			if !ok {
				continue
			}
			for i := range tp.NumEmbeddeds() {
				fmt.Printf("embedded: %s\n", reflect.TypeOf(tp.EmbeddedType(i)))
			}
		}
		if ident.Name == "j" {
			tp, ok := obj.Type().(*types.Alias)
			if !ok {
				continue
			}
			fmt.Println(tp.Obj().IsAlias())
		}
		if ident.Name == "h" {
			tp, ok := obj.Type().(*types.Signature)
			if !ok {
				continue
			}
			fmt.Println(tp.Params().At(0).Name())
			fmt.Println(reflect.TypeOf(tp.Results()))
		}

		switch tp := obj.Type().(type) {
		case *types.Basic:
			fmt.Printf("type of '%s' is %s (types.Basic)\n", ident.Name, tp.String())
		case *types.Pointer:
			fmt.Printf("type of '%s' is %s (types.Pointer)\n", ident.Name, tp.String())
		case *types.Array:
			fmt.Printf("type of '%s' is %s (types.Array)\n", ident.Name, tp.String())
		case *types.Slice:
			fmt.Printf("type of '%s' is %s (types.Slice)\n", ident.Name, tp.String())
		case *types.Map:
			fmt.Printf("type of '%s' is %s (types.Map)\n", ident.Name, tp.String())
		case *types.Chan:
			fmt.Printf("type of '%s' is %s (types.Chan)\n", ident.Name, tp.String())
		case *types.Struct:
			fmt.Printf("type of '%s' is %s (types.Struct)\n", ident.Name, tp.String())
		case *types.Tuple:
			fmt.Printf("type of '%s' is %s (types.Tuple)\n", ident.Name, tp.String())
		case *types.Signature:
			fmt.Printf("type of '%s' is %s (types.Signature)\n", ident.Name, tp.String())
		case *types.Alias:
			fmt.Printf("type of '%s' is %s (types.Alias)\n", ident.Name, tp.String())
		case *types.Named:
			fmt.Printf("type of '%s' is %s (types.Named)\n", ident.Name, tp.String())
		case *types.Interface:
			fmt.Printf("type of '%s' is %s (types.Interface)\n", ident.Name, tp.String())
		case *types.Union:
			fmt.Printf("type of '%s' is %s (types.Union)\n", ident.Name, tp.String())
		case *types.TypeParam:
			fmt.Printf("type of '%s' is %s (types.TypeParam)\n", ident.Name, tp.String())
		}
	}
}
