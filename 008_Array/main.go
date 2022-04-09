package main

import "fmt"

func main() {
	strings := [3]string{"a", "b", "c"}
	fmt.Println(strings)
	fmt.Println(strings[0])
	fmt.Println(len(strings))

	var numbers [2]int8
	numbers[0] = 1
	numbers[1] = 2
	fmt.Println(numbers)

	numbers[0] = 3
	fmt.Println(numbers)

	floats := [2]float32{}
	fmt.Println(floats)

	matrix := [2][2]uint8{{1, 2}, {3, 4}}
	fmt.Println(matrix)
	fmt.Println(matrix[0][1])
}
