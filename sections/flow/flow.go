package flow

import (
	"fmt"
)

func get10() int {
	return 10
}

// ---------------------------------------------------------
// If Statement
// ---------------------------------------------------------
func ifStatement() {
	// If statement
	if true {
		fmt.Println("if statement")
	}

	// If statement with initialization
	if x := 10; x > 0 {
		fmt.Println("if statement with initialization")
	}

	// If-else statement
	if false {
		fmt.Println("if-else statement")
	} else {
		fmt.Println("if-else statement")
	}

	fmt.Println()
}

// ---------------------------------------------------------
// Switch Statement
// ---------------------------------------------------------
func switchStatement() {
	// Switch statement
	switch {
	case true:
		fmt.Println("switch statement")
	}

	// Switch statement with initialization
	switch x := 10; x {
	case 10:
		fmt.Println("switch statement with initialization")
	}

	// Switch statement with multiple cases
	switch x := 10; x {
	case 10, 20:
		fmt.Println("switch statement with multiple cases")
	}

	// Switch statement with default case
	switch x := 10; x {
	case 10:
		fmt.Println("switch statement with default case")
	default:
		fmt.Println("switch statement with default case (default case)")
	}

	// Switch statement with fallthrough and without condition (true by default)
	// fallthrough will execute the next case even if it doesn't match the expression
	switch x := 10; {
	case x < 11:
		fmt.Println("switch statement with fallthrough")
		fallthrough
	case x < 100:
		fmt.Println("switch next case after fallthrough")
	}

	// Switch cases evaluate cases from top to bottom, stopping when a case succeeds
	// case values may be single values or lists of values separated by commas
	// case value can be an expression of any type
	v := 10
	switch v {
	case 1:
		fmt.Println("case 1 (won't be executed)")
	case 9, get10():
		fmt.Println("switch case 9 or 10")
	}

	fmt.Println()
}

// ---------------------------------------------------------
// For Statement
// ---------------------------------------------------------
func forStatement() {
	// Regular for loop
	for i := 0; i < 1; i++ {
		fmt.Println("for statement")
	}

	// while loop analog
	sum := 1
	for sum < 100 {
		sum += sum
	}
	fmt.Println("for in while-like style")

	// Infinite loop
	for {
		fmt.Println("for infinite loop")
		// Loop can be terminated with `break` or `return`
		break
	}

	// For statement with range
	for _, v := range []int{1, 2, 3} {
		fmt.Printf("for statement with range: %v\n", v)
	}

	fmt.Println()
}

// ---------------------------------------------------------
// Defer Statement
// ---------------------------------------------------------
func deferStatement() {
	// Defer statement
	// Defer statement defers the execution of a function until the surrounding function returns.
	// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
	defer fmt.Println("defer top statement: do the clean ups here")

	// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.
	for i := 0; i < 3; i++ {
		defer fmt.Printf("defer statement in loop: %v\n", i)
	}

	fmt.Println()
}

func deferStatementReturn2() (i int) {
	// Deferred function may read and assign to the returning function's named return values.
	// In this case, before the actual function return, anonymous function will be executed and will increment the returning value by 1.
	defer func() { i++ }()

	fmt.Println()

	return 1
}

// ---------------------------------------------------------
// Panic and Recover
// ---------------------------------------------------------
func panicAndRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in panicAndRecover()", r)
		}
	}()

	fmt.Println("panicAndRecover()")
	tryPanic(0)
	fmt.Println("panicAndRecover(): normal exit")
}

func tryPanic(i int) {
	if i > 2 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in tryPanic()", i)
	fmt.Println("Printing in tryPanic()", i)
	tryPanic(i + 1)
}

// ---------------------------------------------------------
// Goto Statement
// ---------------------------------------------------------
func gotoStatement() {
	// Goto statement
	// The goto statement transfers control to the statement with the corresponding label within the same function.
	// The target label must be in the same function and must not be defined in a block.
	// The goto statement cannot be used to jump over variable declarations to the target label.
	goto label
	// Code here is unreachable
	fmt.Println("goto statement")
label:
	fmt.Println("goto statement label")
}

// ---------------------------------------------------------
// Show
// ---------------------------------------------------------
func Show() {
	fmt.Println("======================")
	fmt.Println("/sections/flow/flow.go")
	fmt.Println("======================")

	ifStatement()
	switchStatement()
	forStatement()
	gotoStatement()
	deferStatement()
	deferStatementReturn2()
	panicAndRecover()
}
