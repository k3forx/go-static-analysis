package main

import (
	"fmt"

	"github.com/k3forx/go-static-analysis/types/identifier_resolution"
	"github.com/k3forx/go-static-analysis/types/info_defs"
	"github.com/k3forx/go-static-analysis/types/info_uses"
	"github.com/k3forx/go-static-analysis/types/object"
	pkg "github.com/k3forx/go-static-analysis/types/package"
	"github.com/k3forx/go-static-analysis/types/scope"
	"github.com/k3forx/go-static-analysis/types/type_interface"
)

func main() {
	printDivider("Identifier Resolution")
	identifier_resolution.Do()

	printDivider("Info.Defs")
	info_defs.Do()

	printDivider("Info.Uses")
	info_uses.Do()

	printDivider("Package")
	pkg.Do()

	printDivider("Scope")
	scope.Do()

	printDivider("Object")
	object.Do()

	printDivider("Type Interface")
	type_interface.Do()
}

func printDivider(name string) {
	fmt.Printf("---------------------- %s --------------------------\n", name)
}
