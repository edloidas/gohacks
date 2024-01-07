package main

import (
	"fmt"

	// Importing local package, `gohacks` is a module name defined in `go.mod`
	"gohacks/sections/variables"
)

func main() {
	fmt.Println("Hello, World!")
	variables.Show()
}
