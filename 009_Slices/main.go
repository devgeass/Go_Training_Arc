package main

import "fmt"

func main() {
	slice1 := make([]uint8, 3)
	slice1[0] = 0
	slice1[1] = 1
	slice1[2] = 2
	fmt.Println(slice1)
	// error
	// slice1[3] = 3
	// but you can append to a slice
	slice1 = append(slice1, 3)
	fmt.Println(slice1)
	slice1 = append(slice1, 4, 5)
	fmt.Println(slice1)
	// len function works
	fmt.Println(len(slice1))

	// There is also slice operator :
	// Helps you access elements based on indexes

	// fetch elements index 1 (not 0) included and index 3-1=2
	// slice[low:high]; low included, high not included
	fmt.Println(slice1[1:3])
	// all elements from index 3 onwards (3 included)
	fmt.Println(slice1[3:])
	// all elemnts before index 4 (4 excluded)
	fmt.Println(slice1[:4])

	slice2 := []string{"a", "b", "c"}
	fmt.Println(slice2)
	slice2 = append(slice2, "d")
	fmt.Println(slice2)

	// slices may be multidimentional
	slice3 := [][]uint8{{1}, {1, 2}, {1, 2, 3}}
	fmt.Println(slice3)

}
