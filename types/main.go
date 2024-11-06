package main

import (
	"github.com/k3forx/go-static-analysis/types/identifier_resolution"
	"github.com/k3forx/go-static-analysis/types/info_defs"
	"github.com/k3forx/go-static-analysis/types/info_uses"
	pkg "github.com/k3forx/go-static-analysis/types/package"
)

func main() {
	identifier_resolution.Do()
	printDivider()

	info_defs.Do()
	printDivider()

	info_uses.Do()
	printDivider()

	pkg.Do()
	printDivider()
}

func printDivider() {
	println("-------------------------------------------------")
}
