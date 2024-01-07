// Sample module with examples of Go variables usage and initializations
package functions

import (
	"fmt"
	"log"
	"math/rand"
)

// ---------------------------------------------------------
// Regular Functions
// ---------------------------------------------------------
// Functions available only in this package
func d20() int {
	return rand.Intn(20) + 1
}

// ---------------------------------------------------------
// Multiple Arguments and Return Values
// ---------------------------------------------------------

// Accepts multiple arguments of the same type
// Returns result and error, that can be `nil`
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Variadic Function, that accepts multiple arguments of the same type
func sum(nums ...int) int {
	total := 0
	// Ignore the first return value (index) via `_` (blank identifier)
	for _, n := range nums {
		total += n
	}
	return total
}

// Function with named return value result (of type int)
// Inside the function, result is treated like regular variable. You can assign values to it.
// A bare return statement is used to return the current values of result and err.
// You can have multiple names return values, but cannot mix named and unnamed (regular) return values.
func add(a, b int) (result int) {
	result = a + b
	return
}

// ---------------------------------------------------------
// Anonymous Functions, Closures and Functions as Types
// ---------------------------------------------------------
// Define a function type
type Adder func(int) int

// Anonymous function that returns a function (of type Adder)
var adder = func() Adder {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

// ---------------------------------------------------------
// Functions as Arguments
// ---------------------------------------------------------
func doOperation(v int, op func(int) int) int {
	return op(v)
}

// ---------------------------------------------------------
// Show
// ---------------------------------------------------------
// Exported function (available outside of this package) start with the capital letter
func Show() {
	fmt.Println("================================")
	fmt.Println("/sections/functions/functions.go")
	fmt.Println("================================")

	fmt.Printf("d20: %v\n", d20())

	division, err := divide(10, 2)
	if err != nil {
		// Log error and exit program with non-zero exit code
		log.Fatal(err)
	} else {
		fmt.Printf("Division: %v\n", division)
	}

	fmt.Printf("Sum: %v\n", sum(1, 2, 3, 4, 5))

	fmt.Printf("Add: %v\n", add(1, 2))

	a := adder()
	fmt.Printf("Added via Adder: (+2): %v, (+3): %v\n", a(2), a(3))

	fmt.Printf("Do Operation: %v\n", doOperation(10, a))

	fmt.Println()
}
