package main

import "fmt"

func main() {
	var a uint8 = 45
	var b uint8 = 5
	// var c float32 = 2.5

	fmt.Println("**** Arithmetic Operators ****")
	// Addition
	fmt.Printf("%v + %v = %v\n", a, b, a+b)
	// Subtraction
	fmt.Printf("%v - %v = %v\n", a, b, a-b)
	// Multiplication
	fmt.Printf("%v * %v = %v\n", a, b, a*b)
	// Division
	fmt.Printf("%v / %v = %v\n", a, b, a/b)
	// Modulo
	fmt.Printf("%v %% %v = %v\n", a, b, a%b)
	// Increment
	fmt.Printf("%v++ = ", a)
	a++
	fmt.Printf("%v\n", a)
	// Decrement
	fmt.Printf("%v-- = ", b)
	b--
	fmt.Printf("%v\n", b)

	// Restore
	a--
	b++

	// String addition i.e. concatenation
	fmt.Println("Attack" + " " + "on" + " " + "Titan")

	fmt.Println("**** Relational Operators ****")
	// is equal to?
	fmt.Printf("%v == %v? %v\n", a, b, a == b)
	// is not equal to?
	fmt.Printf("%v != %v? %v\n", a, b, a != b)
	// is less than
	fmt.Printf("%v < %v? %v\n", a, b, a < b)
	// is greater than
	fmt.Printf("%v > %v? %v\n", a, b, a > b)
	// is less than or equal to
	fmt.Printf("%v <= %v? %v\n", a, b, a <= b)
	// is greater than or equal to
	fmt.Printf("%v >= %v? %v\n", a, b, a >= b)

	fmt.Println("**** Logical Operators ****")
	t := true
	f := false

	// logical AND
	fmt.Printf("%v && %v = %v\n", t, f, t && f)
	// logical OR
	fmt.Printf("%v || %v = %v\n", t, f, t || f)
	// logical NOT
	fmt.Printf("!%v = %v\n", t, !t)
	fmt.Printf("!%v = %v\n", f, !f)

	fmt.Println("**** Bitwise Operators ****")
	var c uint8 = 7
	var d uint8 = 10
	// bitwise AND
	fmt.Printf("%v & %v = %v\n", c, d, c&d)
	// bitwise OR
	fmt.Printf("%v | %v = %v\n", c, d, c|d)
	// bitwise XOR
	fmt.Printf("%v ^ %v = %v\n", c, d, c^d)
	// bitwise LEFT SHIFT
	fmt.Printf("%v << %v = %v\n", c, 2, c<<2)
	// bitwise RIGHT SHIFT
	fmt.Printf("%v >> %v = %v\n", c, 2, c>>2)

}
