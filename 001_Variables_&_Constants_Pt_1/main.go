package main

import "fmt"

var d1 int    // declaration with type
var d2 = 60   // declaration with initialization (type inferred)
const f1 = 70 // declaration with initialization (type inferred)
// const f2 - error, constants must be initialized
// d3 := 80 - not allowed short hand outside functions

func main() {
	var v1 = 10       // declaration with initialization (type inferred)
	var v2 int = 20   // declared and initialized with type
	var v3 int        // declaration with type
	v3 = 30           // initialization after declaration
	v4 := 40          // short hand declaration and initialization
	d1 = 50           // declare outside, initialize inside
	const f2 int = 80 // declared and initialized with type

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(v3)
	fmt.Println(v4)
	fmt.Println(d1)
	fmt.Println(d2)
	fmt.Println(f1)
	fmt.Println(f2)

	// v3 = "Hello" - cannot reassign type, specified earlier
	// v1 = "Hello" - cannot reassign type, inferred earlier
	// f1 = 80 - cannot reassign value of a constant
}
