package main

import "fmt"

func add_sub(a, b int16) (int16, int16) {
	return a + b, a - b
}

func summation(numbers ...int) {
	fmt.Printf("%v -> sum: ", numbers)
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("%v\n", sum)
}

func main() {
	sum, difference := add_sub(10, 20)

	fmt.Printf("Sum: %v\nDifference: %v\n", sum, difference)

	_, diff := add_sub(20, 10)
	fmt.Printf("Diff: %v\n", diff)

	summation(1, 4)
	summation(1, 6, 10, 30)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	summation(arr...)
}
