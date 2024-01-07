// Sample module with examples of Go variables usage and initializations
package variables

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

// Constant (cannot be changed)
const Pi = 3.14

// ---------------------------------------------------------
// Function Level Variables
// ---------------------------------------------------------
func ExportedFunction() int {
	// Variables declared in function will be scoped to the function

	// Default syntax
	var localInt int // 0

	// Second return value is dropped
	x, _ := unexportedFunction(1.0, 2.0)

	// Type conversion
	localInt = int(x)

	return localInt
}

func unexportedFunction(argX, argY float64) (float64, float64) {
	// Constant
	const Multiplier = 2

	// Short variable declaration
	x, y := argX, argY // float inferred

	// Assignment to new value from local scope
	x = Multiplier*x + gX
	y += gY

	return x, y
}
