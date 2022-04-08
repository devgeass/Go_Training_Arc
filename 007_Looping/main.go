package main

import "fmt"

func main() {
	// traditional loop
	for i := 0; i <= 9; i++ {
		fmt.Printf("%v ", i)
	}
	fmt.Println()

	// loop with condition only
	k := 1
	for k <= 3 {
		fmt.Printf("%v ", k)
		k++
	}
	fmt.Println()

	// infinite loop
	// for {
	// any code here
	// }

	// continue statement
	for j := 1; j <= 10; j++ {
		if j%2 == 0 {
			continue
		}
		fmt.Printf("%v ", j)
	}
	fmt.Println()

	// break statement
	for j := 1; j <= 10; j++ {
		if j == 3 {
			break
		}
		fmt.Printf("%v ", j)
	}
	fmt.Println()
}
