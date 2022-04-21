package main

import "fmt"

func main() {
	cities := []string{"New York", "London", "Toronto", "Los Angeles"}

	for index, city := range cities {
		fmt.Printf("%v: %v\n", index, city)
	}

	// ignore the index
	for _, city := range cities {
		fmt.Println(city)
	}

	fruits := map[string]string{"a": "Apple", "b": "Banana", "c": "Cherry"}

	for key, value := range fruits {
		fmt.Printf("%v -> %v\n", key, value)
	}

	for key := range fruits {
		fmt.Println(key)
	}
}
