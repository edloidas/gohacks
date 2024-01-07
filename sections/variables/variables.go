// Sample module with examples of Go variables usage and initializations
package variables

import (
	"fmt"
)

// ---------------------------------------------------------
// Package Level (Global) Variables
// ---------------------------------------------------------
// Explicit type declaration, variables initialized with default values
var globalString string // ""
var globalInt int       // 0
var globalBool bool     // false

// Implicit type declaration
var globalStringHello = "Hello, World!" // string inferred

// Multiple variable declaration
var gX, gY = 1.0, float64(2)
var gA, gB string = "A", "B"

// Grouped variable declaration
var (
	gInt   int = 10
	gZ         = 3.0
	gC, gD string
)

// Exported (starts with capital letter) constant (defined with `const` keyword, cannot be changed)
const Pi = 3.14

// ---------------------------------------------------------
// Function Level Variables
// ---------------------------------------------------------
func Show() {
	fmt.Println("================================")
	fmt.Println("/sections/variables/variables.go")
	fmt.Println("================================")

	// Variables declared in function will be scoped to the function

	// Constant
	const Multiplier = 2

	// Default syntax
	var localInt int // 0

	// Short variable declaration
	x, y := gX, gY // float inferred

	// Assignment to new value from local scope
	x = Multiplier*x + gX
	y += gY

	// Type conversion
	localInt = int(x)

	fmt.Printf("globalString: %v, globalInt: %v, globalBool: %v\n", globalString, globalInt, globalBool)
	fmt.Printf("globalStringHello: %v\n", globalStringHello)
	fmt.Printf("gX: %v, gY: %v, gA: %v, gB: %v, gZ: %v, gC: %v, gD: %v\n", gX, gY, gA, gB, gZ, gC, gD)
	fmt.Printf("localInt: %v, x: %v, y: %v\n", localInt, x, y)
	fmt.Println()
}
