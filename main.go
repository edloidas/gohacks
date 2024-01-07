package main

import (

	// Importing local package, `github.com/edloidas/gohacks` is a module name defined in `go.mod`
	"github.com/edloidas/gohacks/sections/flow"
	"github.com/edloidas/gohacks/sections/functions"
	"github.com/edloidas/gohacks/sections/variables"
)

func main() {
	variables.Show()
	functions.Show()
	flow.Show()
}
