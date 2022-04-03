package main

import "fmt"

var d1 int  // declared, not initialized and unused, no error
var d2 = 50 // unused, no error

const f1 = 60 // unused, no error

func main() {
	// var m1 int - throws error because it is unused
	// m2 := 40 - throws error because it is unused
	d2 := 60              // declaring d2 in function scope takes preference
	var h1, h2 = 1, "Wow" // multiple declaration and initialization (type inferred)
	var k1, k2 int = 2, 3 // multiple declaration and initialization with type
	var (
		c1             = 20 // type inferred
		c2 int         = 30 // type specified
		c3 int              // not initialized
		c4 = "Really?"      // type inferred
	)
	c3 = 40
	const f2 int = 70 // unused but still no error
	fmt.Println(d2)
	fmt.Println(h1)
	fmt.Println(h2)
	fmt.Println(k1)
	fmt.Println(k2)
	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
}
