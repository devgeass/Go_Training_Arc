package main

import "fmt"

func main() {
	var m int = 5
	// n is a pointer that stores address of m
	var n *int = &m

	// value of m
	fmt.Println("m =", m)
	// address of m
	fmt.Println("&m =", &m)
	// value of n which is nothing but address of m
	fmt.Println("n=", n)
	// value stored in address stored in n -> value of m
	fmt.Println("*n=", *n)
	// address of n
	fmt.Println("&n=", &n)

	// deference the reference of m -> value of m
	fmt.Println("*(&m) =", *(&m))

	arr := [3]int{1, 2, 3}
	// copy arr
	arr_copy := arr
	// modify copy arr
	arr_copy[0] = 10
	fmt.Println("arr_copy =", arr_copy)
	// no change in original arr
	fmt.Println("arr =", arr)
	// copy arr to pointer variable
	// var arr_copy_ptr *[3]int = &arr
	arr_copy_ptr := &arr
	// modify pointer type arr value
	arr_copy_ptr[0] = 20
	// print value stored in address stored by arr_copy_ptr -> arr
	fmt.Println("arr_copy_ptr =", *arr_copy_ptr)
	// print arr, it is modified
	fmt.Println("arr =", arr)

	// Pointer Chaining
	i := "devgeass"
	// j stores address of i
	j := &i
	// k stores address of j
	k := &j
	fmt.Println("i =", i)
	fmt.Println("&i =", &i)
	fmt.Println("j =", j)
	fmt.Println("*j =", *j)
	fmt.Println("&j =", &j)
	fmt.Println("k =", k)
	fmt.Println("*k =", *k)
	fmt.Println("**k =", **k)
}
