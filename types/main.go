package main

import (
	"fmt"

	"github.com/k3forx/go-static-analysis/types/type_interface"
)

func main() {
	// printDivider("Identifier Resolution")
	// identifier_resolution.Do()

	// printDivider("Info.Defs")
	// info_defs.Do()

	// printDivider("Info.Uses")
	// info_uses.Do()

	// printDivider("Package")
	// pkg.Do()

	// printDivider("Scope")
	// scope.Do()

	// printDivider("Object")
	// object.Do()

	printDivider("Type Interface")
	type_interface.Do()

	// printDivider("Info.Selections")
	// info_selections.Do()

	// printDivider("Config")
	// config.Do()
}

func printDivider(name string) {
	fmt.Printf("---------------------- %s --------------------------\n", name)
}
