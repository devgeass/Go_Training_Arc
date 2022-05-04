package main

import "fmt"

func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}

func main() {
	add5 := makeAdder(5)
	fmt.Println(add5(4))
	fmt.Println(add5(10))
	add10 := makeAdder(10)
	fmt.Println(add10(4))
	fmt.Println(add10(10))
}
