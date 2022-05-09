package main

import (
	"fmt"
)

func main() {
	// array
	arr := [5]uint8{1, 2, 3, 4, 5}

	fmt.Println("arr =", arr)
	fmt.Println("len(arr) =", len(arr))
	fmt.Println("Shorthand Array Declartion of arr2")
	arr2 := [...]uint8{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr2 =", arr2)
	fmt.Println("Copying arr2 to arr3")
	arr3 := arr2
	fmt.Println("Modifying arr3 first element")
	arr3[0] = 10
	fmt.Println("arr3 =", arr3)
	fmt.Println("arr2 remains unchanged")
	fmt.Println("arr2 =", arr2)

	// slices
	slc := []uint8{1, 2, 3, 4, 5}

	fmt.Println("slc =", slc)
	fmt.Println("Appending new elements to slc")
	slc = append(slc, 6, 7, 8, 9)
	fmt.Println("slc =", slc)
	slc2 := []uint8{10, 11}
	fmt.Println("Appending to slc from slc2 using spread operator")
	slc = append(slc, slc2...)
	fmt.Println("slc =", slc)
	fmt.Println("len(slc) =", len(slc))
	fmt.Println("Copying slc to slc3")
	slc3 := slc
	fmt.Println("Modifying slc3 first element")
	slc3[0] = 10
	fmt.Println("slc3 =", slc3)
	fmt.Println("slc ALSO CHANGES")
	fmt.Println("slc =", slc)

	// maps
	map1 := map[string]uint8{"a": 1, "b": 2, "c": 3}

	fmt.Println("map1 =", map1)
	fmt.Println("len(map1) =", len(map1))
	fmt.Println("Copying map1 to map2")
	map2 := map1
	fmt.Println("Modifying key \"a\" of map2")
	map2["a"] = 4
	fmt.Println("map2 =", map2)
	fmt.Println("map1 ALSO CHANGES")
	fmt.Println("map1 =", map1)
	fmt.Println("Deleting key \"a\" of map2")
	delete(map2, "a")
	fmt.Println("map2 =", map2)
	fmt.Println("map1 ALSO CHANGES")
	fmt.Println("map1 =", map1)
}
