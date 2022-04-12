package main

import "fmt"

func main() {
	m := make(map[string]uint)
	m["a"] = 1
	m["b"] = 2
	fmt.Println(m)

	fmt.Println(m["a"])

	fmt.Println(len(m))

	delete(m, "a")

	fmt.Println(m)

	c, isPresent := m["c"]
	if isPresent == true {
		fmt.Println(c)
	} else {
		fmt.Println(isPresent)
	}

	mp := map[string]uint{"a": 1, "b": 2}
	fmt.Println(mp)
}
