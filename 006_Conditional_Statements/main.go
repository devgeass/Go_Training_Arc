package main

import "fmt"

func main() {
	t := true
	f := false

	// if condition is true, execute if-block
	// else, execute else-block
	if t {
		fmt.Println("I am true!")
	} else {
		fmt.Println("I am false!")
	}

	// You can have if-block without an else
	if !f {
		fmt.Println("I am not false!")
	}

	// for multiple conditions, use if-else if-else ladder
	a := 0
	if a > 0 {
		fmt.Println("Greater than 0!")
	} else if a < 0 {
		fmt.Println("Less than 0!")
	} else {
		fmt.Println("It is 0!")
	}

	// You write a statement scoped within if block
	// the statement precedes the condition
	if n := 3; n%2 == 0 {
		fmt.Println("Number is even!")
	} else {
		fmt.Println("Number is odd!")
	}

	// n is not scoped outside the conditional block
	// fmt.Println(n) throws error

	b := 2
	// single case switch
	switch b {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("Default")
	}

	// multi-case switch
	switch b {
	case 1, 2, 3:
		fmt.Println("In (1,2,3)")
	case 4, 5, 6:
		fmt.Println("In (4,5,6)")
	case 7, 8, 9:
		fmt.Println("In (7,8,9)")
	default:
		fmt.Println("Above 9")
	}
}
