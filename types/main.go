package main

import (
	"github.com/k3forx/go-static-analysis/types/identifier_resolution"
	"github.com/k3forx/go-static-analysis/types/info_uses_defs"
)

func main() {
	println()
	identifier_resolution.Do()

	println()
	info_uses_defs.Do()
}
