package main

import "fmt"

func greet() {
	fmt.Println("Have a nice day!")
}

func greetByName(name string) {
	fmt.Printf("Have a nice day, %v!\n", name)
}

func nani() string {
	return "Omae wa mou shindeiru!"
}

func add(a, b uint8) uint8 {
	return a + b
}

func main() {
	greet()
	greetByName("Thomas")
	fmt.Println(nani())
	fmt.Println(add(1, 2))
}
