package scope

// const pkgStr = `package main

// var a int

// func main() {
// 	b := "hello"
// 	_ = b
// }
// `

func Do() {
	// fmt.Printf("Universe's element names: %+v\n", types.Universe.String()

	// boolObj := universe.Lookup("bool")
	// obj := boolObj.(*types.TypeName)
	// fmt.Println(obj.Parent().String())

	// fset := token.NewFileSet()
	// f, err := parser.ParseFile(fset, "main.go", pkgStr, parser.Mode(0))
	// if err != nil {
	// 	log.Fatalf("parsing file: %v", err)
	// }

	// conf := types.Config{}
	// pkgInfo, err := conf.Check("main.go", fset, []*ast.File{f}, nil)
	// if err != nil {
	// 	log.Fatalf("checking types: %v", err)
	// }

	// fmt.Printf("Name: %s\n", pkgInfo.Name())
	// fmt.Printf("# of scope in %s: %+v\n", pkgInfo.Name(), pkgInfo.Scope().NumChildren())

	// for _, name := range pkgInfo.Scope().Names() {
	// 	obj := pkgInfo.Scope().Lookup(name)
	// 	switch tp := obj.(type) {
	// 	case *types.Var:
	// 		fmt.Printf("scope: %+v\n", tp.String())
	// 	}
	// }

	// for i := range pkgInfo.Scope().NumChildren() {
	// 	child := pkgInfo.Scope().Child(i)
	// 	for _, name := range child.Names() {
	// 		obj := child.Lookup(name)
	// 		fmt.Printf("child obj: %+v\n", obj)
	// 	}
	// }
}
