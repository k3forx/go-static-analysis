# go/typesパッケージ

## 概要

Goの型チェックは以下の3つのことをする。

1. identifier resolution
    - プログラムにあるすべての"名前"について、その名前が参照する識別子を特定する
2. type deduction
    - プログラムにある"各式"について、その式が持つ型を特定し、式に型がない場合や、その文脈に対して不適切な型である場合にはエラーを返す
3. constant evaluation
    - プログラムにあるすべての"定数表現"

## Identifier Resolution

- identifier resolutionによって `ast.Ident` と `Object` のマッピングが作られる

## Example

## `Package` 構造体

構造体の詳細

```bash
❯ go doc go/types.Package
package types // import "go/types"

type Package struct {
        // Has unexported fields.
}
    A Package describes a Go package.

var Unsafe *Package
func NewPackage(path, name string) *Package
func (pkg *Package) Complete() bool
func (pkg *Package) GoVersion() string
func (pkg *Package) Imports() []*Package
func (pkg *Package) MarkComplete()
func (pkg *Package) Name() string
func (pkg *Package) Path() string
func (pkg *Package) Scope() *Scope
func (pkg *Package) SetImports(list []*Package)
func (pkg *Package) SetName(name string)
func (pkg *Package) String() string
```

## `Scope` 構造体

まとめ

- オブジェクトの集合とそれらを含んでいる (親の) スコープと含んだ (子の) パッケージのスコープを表現する
- 名前とObjectのマッピング

packageの字句ブロック (lexical block) を保持する構造体。packageレベルで定義されている名前付きのentityとobjectにアクセスできる。

```bash
❯ go doc go/types.Scope
package types // import "go/types"

type Scope struct {
        // Has unexported fields.
}
    A Scope maintains a set of objects and links to its containing (parent) and
    contained (children) scopes. Objects may be inserted and looked up by name.
    The zero value for Scope is a ready-to-use empty scope.

var Universe *Scope
func NewScope(parent *Scope, pos, end token.Pos, comment string) *Scope
func (s *Scope) Child(i int) *Scope
func (s *Scope) Contains(pos token.Pos) bool
func (s *Scope) End() token.Pos
func (s *Scope) Innermost(pos token.Pos) *Scope
func (s *Scope) Insert(obj Object) Object
func (s *Scope) Len() int
func (s *Scope) Lookup(name string) Object
func (s *Scope) LookupParent(name string, pos token.Pos) (*Scope, Object)
func (s *Scope) Names() []string
func (s *Scope) NumChildren() int
func (s *Scope) Parent() *Scope
func (s *Scope) Pos() token.Pos
func (s *Scope) String() string
func (s *Scope) WriteTo(w io.Writer, n int, recurse bool)
```

`Names` メソッドは名前の集合を返す。`Lookup` メソッドは与えられた名前に対するobjectをお返す。

## オブジェクト

- identifier resolutionのタスクは `ast.Ident` を *object* にマップすること。
- `1` を

```bash
❯ go doc go/types.Object
package types // import "go/types"

type Object interface {
        Parent() *Scope // scope in which this object is declared; nil for methods and struct fields
        Pos() token.Pos // position of object identifier in declaration
        Pkg() *Package  // package to which this object belongs; nil for labels and objects in the Universe scope
        Name() string   // package local object name
        Type() Type     // object type
        Exported() bool // reports whether the name starts with a capital letter
        Id() string     // object name if exported, qualified name if not exported (see func Id)

        // String returns a human-readable string of the object.
        String() string

        // Has unexported methods.
}
    An Object describes a named language entity such as a package, constant,
    type, variable, function (incl. methods), or label. All objects implement
    the Object interface.

func LookupFieldOrMethod(T Type, addressable bool, pkg *Package, name string) (obj Object, index []int, indirect bool)
```

`ast.Ident` にパッケージの情報や型の情報、Scopeの情報を肉付けしたもの？

`object` インターフェイスを満たす構造体は以下の8つ。

```bash
Object = *Func         // function, concrete method, or abstract method
       | *Var          // variable, parameter, result, or struct field
       | *Const        // constant
       | *TypeName     // type name
       | *Label        // statement label
       | *PkgName      // package name, e.g. json after import "encoding/json"
       | *Builtin      // predeclared function such as append or len
       | *Nil          // predeclared nil
```

## Identifier Resolution

識別子とObjectの関係は `Check` 関数の最後の引数に渡す `types.Info` 構造体に保存される。APIを呼ぶ側で何の情報が必要かを制御することができるようになっている。

```bash
type Info struct {
	Defs       map[*ast.Ident]Object
	Uses       map[*ast.Ident]Object
	Implicits  map[ast.Node]Object
	Selections map[*ast.SelectorExpr]*Selection
	Scopes     map[ast.Node]*Scope
	...
}
```

`map[*ast.Indent]Object` 型で最も重要な2つのフィールドは

- `Defs`: 識別子が定義されている箇所を保持する
- `Uses`: 識別子が参照されている箇所を保持する
- `Implicits`: 
- `Selections`: 
- 

以下のコードで試してみる。

```bash
func PrintDefsUses(fset *token.FileSet, files ...*ast.File) error {
	conf := types.Config{Importer: importer.Default()}
	info := &types.Info{
		Defs: make(map[*ast.Ident]types.Object),
		Uses: make(map[*ast.Ident]types.Object),
	}
	_, err := conf.Check("hello", fset, files, info)
	if err != nil {
		return err // type error
	}

	for id, obj := range info.Defs {
		fmt.Printf("%s: %q defines %v\n",
			fset.Position(id.Pos()), id.Name, obj)
	}
	for id, obj := range info.Uses {
		fmt.Printf("%s: %q uses %v\n",
			fset.Position(id.Pos()), id.Name, obj)
	}
	return nil
}
```

## `Scope` 構造体

- 名前からobjectへのマップを持っている

```go
type Scope struct{ ... }

func (s *Scope) Names() []string
func (s *Scope) Lookup(name string) Object
```

`Names` でマッピング内の名前のセットをソートして返す。`Lookup` を使えば、名前に対応する `Object` を探せる。

```go
for _, name := range scope.Names() {
	fmt.Println(scope.Lookup(name))
}
```

`go/types` パッケージのscopeは字句スコープ (静的スコープ) を表す。字句スコープは字句環境で構成される。具体例として以下のコードを考えてみる。

```go
package main

import "fmt"

func main() {
	const message = "hello, world"
	fmt.Println(message)
}
```

上記のプログラムには4つの字句ブロックが存在する。

- universalブロック: 予約語がobjectにマップされている
- packageブロック: 

```go
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

const hello = `package main

import "fmt"

func main() {
	const message = "hello, world"
	fmt.Println(message)
}
`

func main() {
	fset := token.NewFileSet()

	f, err := parser.ParseFile(fset, "hello.go", hello, 0)
	if err != nil {
		log.Fatal(err)
	}

	// 型チェックのための設定を行う
	// Importerフィールドはimportの情報を解析するための設定
	conf := types.Config{Importer: importer.Default()}

	// Check関数はtypes.Configに基づいて型チェックを行う
	// 返り値はPackage型になる
	// 最後の引数のtypes.Infoにast.IdentとObjectの紐付けの結果が格納される

	info := &types.Info{
		Scopes: map[ast.Node]*types.Scope{},
	}

	pkg, err := conf.Check("cmd/hello", fset, []*ast.File{f}, info)
	if err != nil {
		log.Fatal(err)
	}

	scope := pkg.Scope()
	for _, name := range scope.Names() {
		fmt.Println(scope.Lookup(name))
	}
}
```

- 字句環境をlookupする
  - `LookupParent` メソッドを使う

```go
package main

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"log"
	"strings"
)

const hello = `
package main

import "fmt"

// append
func main() {
	// fmt
	fmt.Println("Hello, world")
	// main
	main, x := 1, 2
	// main
	print(main, x)
	// x
}
// x
`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "hello.go", hello, parser.ParseComments)
	if err != nil {
		log.Fatal(err)
	}

	conf := types.Config{Importer: importer.Default()}
	pkg, err := conf.Check("cmd/hello", fset, []*ast.File{f}, nil)
	if err != nil {
		log.Fatal(err)
	}

	for _, comment := range f.Comments {
		pos := comment.Pos()
		name := strings.TrimSpace(comment.Text())
		fmt.Printf("At %s, \t%q = ", fset.Position(pos), name)

		inner := pkg.Scope().Innermost(pos)
		if _, obj := inner.LookupParent(name, pos); obj != nil {
			fmt.Println(obj)
		} else {
			fmt.Println("not found")
		}
	}
}
```

実行するとこんな感じになる。

```bash
❯ go run ./types
At hello.go:6:1,        "append" = builtin append
At hello.go:8:2,        "fmt" = package fmt
At hello.go:10:2,       "main" = func cmd/hello.main()
At hello.go:12:2,       "main" = var main int
At hello.go:14:2,       "x" = var x int
At hello.go:16:1,       "x" = not found
```

## 初期化の順序 (Initialization Order)

packageレベルの初期化の順序を保持する。`Info` 構造体の `InitOrder` フィールドに情報が格納される。

```go
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
```

実行すると `var x int = 1` の情報が取れる。

```bash
❯ go run ./types
initOrder: x = 1
Lhs: x, Rhs: 1
```

## 型 (Types)

`Type` はインターフェイス。

```go
type Type interface {
	Underlying() Type
}
```

`Type` インターフェイスを満たす具体な型は以下。

```go
Type = *Basic
     | *Pointer
     | *Array
     | *Slice
     | *Map
     | *Chan
     | *Struct
     | *Tuple
     | *Signature
     | *Named
     | *Interface
```

### 基本的な型 (Basic Types)


```go
❯ go doc go/types.Basic
package types // import "go/types"

type Basic struct {
        // Has unexported fields.
}
    A Basic represents a basic type.

func (b *Basic) Info() BasicInfo
func (b *Basic) Kind() BasicKind
func (b *Basic) Name() string
func (b *Basic) String() string
func (b *Basic) Underlying() Type
```
