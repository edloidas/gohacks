package moretypes

import (
	"fmt"
)

// ---------------------------------------------------------
// Structs
// ---------------------------------------------------------
type Vertex struct {
	X int
	Y int
}

func structs() {
	fmt.Println("Structs")

	v1 := Vertex{1, 2}
	v1.X = 4
	fmt.Printf("v1: %v\n", v1)

	v2 := Vertex{X: 1} // Y:0 is implicit
	fmt.Printf("v2: %v\n", v2)

	v3 := Vertex{} // X:0 and Y:0
	fmt.Printf("v3: %v\n", v3)

	fmt.Println()
}

// ---------------------------------------------------------
// Arrays
// ---------------------------------------------------------
func arrays() {
	fmt.Println("Arrays")

	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Printf("a: %v\n", a)

	primes := [7]int{2, 3, 5, 7, 11, 13} // Value by index 7 will be 0 by default
	fmt.Printf("primes: %v\n", primes)
}

// ---------------------------------------------------------
// Slices
// ---------------------------------------------------------
func slices() {
	fmt.Println("Slices")

	primes := [6]int{2, 3, 5, 7, 11, 13} // Value by index 7 will be 0 by default

	// A slice literal is like an array literal without the length

	// Slice is a reference to the array, and do not store any data
	// Changing elements of a slice modifies the corresponding elements of its underlying array
	var s1 []int = primes[1:4] // [1:4) (including index 1, excluding index 4)
	var s2 = primes[2:5]       // [2:5) (including index 2, excluding index 5), type is inferred
	s1[1] = 100
	fmt.Printf("s1: %v, s2: %v, primes: %v\n", s1, s2, primes)

	// Slice has both a length and a capacity
	s3 := []int{3, 5, 2, 9, 7, 1}

	// When slicing, you may omit the high or low bounds to use their defaults instead
	// The default is zero for the low bound and the length of the slice for the high bound
	fmt.Printf("s3[:2]: %v, s3[2:]: %v, s3[:]: %v\n", s3[:2], s3[2:], s3[:])

	// Slice has both a length and a capacity
	s4 := []int{0, 1, 2, 3, 4, 5}
	printSlice(s4)

	s4 = s4[:0] // Slice the slice to give it zero length
	printSlice(s4)

	s4 = s4[:4] // Extend its length
	printSlice(s4)

	s4 = s4[2:] // Drop its first two values
	printSlice(s4)

	// Capacity of slice after dropping first two values is 4 and cannot be extended by regaining dropped values

	// The zero value of a slice is nil.
	// A nil slice has a length and capacity of 0 and has no underlying array.
	var s5 []int
	printSlice(s5)

	// Slice can contain any type, including other slices
	matrix := [][]int{
		{1, 2},
		{3, 4},
	}
	matrix[0][1] = 100
	fmt.Printf("matrix: %v\n", matrix)

	fmt.Println()
}

func sliceMake() {
	fmt.Println("Slice: Make")
	// The make function allocates a zeroed array and returns a slice that refers to that array
	a := make([]int, 3) // len(a)=3, cap(a)=3
	printSlice(a)

	b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	printSlice(b)

	c := b[:4] // len(c)=4, cap(c)=5
	printSlice(c)

	d := c[2:4] // len(d)=2, cap(d)=2
	printSlice(d)

	matrix := make([][]int, 3)
	for i := range matrix {
		matrix[i] = make([]int, 3)
		for j := range matrix[i] {
			matrix[i][j] = i + j
		}
	}
	fmt.Printf("matrix: %v\n", matrix)

	fmt.Println()
}

func sliceAppend() {
	fmt.Println("Slice: Append")

	var s []int
	printSlice(s)

	// append to nil slice
	s = append(s, 0) // len(s)=1, cap(s)=1
	p := &s
	printSliceN(s, "*p")
	printSlice(*p)

	// The slice grows as needed
	s = append(s, 1, 2) // len(s)=3, cap(s)=3
	printSlice(s)
	printSliceN(s, "*p")

	// We can add more than one element at a time
	s = append(s, 4, 5, 6) // len(s)=7, cap(s)=8
	printSlice(s)

	// We can append a slice to itself
	s = append(s, s...) // len(s)=14, cap(s)=16
	printSlice(s)

	fmt.Println()
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceN(s []int, name string) {
	fmt.Printf("(%v) len=%d cap=%d %v\n", name, len(s), cap(s), s)
}

// ---------------------------------------------------------
// Maps
// ---------------------------------------------------------
func maps() {
	fmt.Println("Maps")

	// New map is created and initialized with map with 2 keys
	location := map[string]Vertex{
		"Seattle": {47, -122},
		"Tokyo":   {35, 140},
	}
	// Add new key-value pair
	location["Barcelona"] = Vertex{41, 2}
	// Delete key-value pair
	delete(location, "Seattle")
	// Update value
	location["Tokyo"] = Vertex{35, 139}

	fmt.Printf("m: %v\n", location)

	_, ok := location["Seattle"]
	fmt.Printf("location[\"Seattle\"]: %v, ok: %v\n", location["Seattle"], ok)

	// Practically the same as above
	// You can also provide the size of the map to preallocate the internal data structure, e.g.:
	// coordinates := make(map[string]Vertex, 10)
	coordinates := make(map[string]Vertex)
	coordinates["center"] = Vertex{0, 0}
	coordinates["top"] = Vertex{0, 100}
	fmt.Printf("coordinates: %v\n", coordinates)
	fmt.Printf("coordinates[\"center\"]: %v\n", coordinates["center"])

	fmt.Println()
}

// ---------------------------------------------------------
// Pointers
// ---------------------------------------------------------
func pointers() {
	fmt.Println("Pointers")

	var x int
	var p *int // p is a pointer to int (nil by default)
	p = &x     // p points to x
	*p = 10    // equivalent to x = 10
	*p++       // equivalent to x++
	fmt.Printf("x: %v, p: %v, *p: %v\n", x, p, *p)

	var pNew = new(string) // p, of type *int, points to an unnamed int variable
	*pNew = "unnamed variable"
	fmt.Printf("pNew: %v, *pNew: %v\n", pNew, *pNew)

	v := Vertex{1, 2}
	pV := &v   // equivalent to pV := &Vertex{1, 2}
	pV.X = 1e9 // equivalent to (*pV).X = 1e9
	fmt.Printf("v: %v, pV: %v, *pV: %v\n", v, pV, *pV)

	fmt.Println()
}

// ---------------------------------------------------------
// Show
// ---------------------------------------------------------
func Show() {
	fmt.Println("================================")
	fmt.Println("/sections/moretypes/moretypes.go")
	fmt.Println("================================")

	structs()
	arrays()
	slices()
	sliceMake()
	sliceAppend()
	maps()
	pointers()
}
