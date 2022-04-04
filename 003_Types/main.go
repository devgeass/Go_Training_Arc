package main

import "fmt"

func main() {
	// Integer Type
	var a uint = 1 // unsigned int
	// Other Unsigned Int Types Include:
	// uint8 : 0 to 255
	// uint16 : 0 to 65535
	// uint32 : 0 to 4294967295
	// uint64 : 0 to 18446744073709551615
	fmt.Printf("a is of type %T\n", a)
	var b int = -16 // signed int
	// Other Unsigned Int Types Include:
	// int8  : -128 to 127
	// int16 :  -32768 to 32767
	// int32 : -2147483648 to 2147483647
	// int64 : -9223372036854775808 to 9223372036854775807
	fmt.Printf("b is of type %T\n", b)

	// Float Type
	var c float32 = 4.65
	var d float64 = 10.75
	fmt.Printf("c is of type %T\n", c)
	fmt.Printf("d is of type %T\n", d)

	// Complex Type
	var e complex64 = 10 + 7i
	var f complex128 = 14.25 - 8.05i
	fmt.Printf("e is of type %T\n", e)
	fmt.Printf("f is of type %T\n", f)

	// String Type
	var g string = "Yo!"
	fmt.Printf("g is of type %T\n", g)

	// Boolean Type
	var h bool = true
	fmt.Printf("h is of type %T\n", h)

	// Character Type
	var i rune = 'a'
	fmt.Printf("i is of type %T\n", i)

	// invalid types uint and int
	// fmt.Println("a + b", a + b) - error
	// Convert to same type
	fmt.Println("a + b = ", int(a)+b)

	// invalid types int and float
	// fmt.Println("b + c", b + c) - error
	// Convert to same type
	fmt.Println("b + c = ", float32(b)+c)

	var o uint8 = 2
	var j int32 = 45
	var k int64 = 78
	var l float32 = 90.25

	fmt.Println(int32(o) + j)
	fmt.Println(int64(j) + k)
	fmt.Println(float32(k) + l)
	fmt.Println(k + int64(l))
}
