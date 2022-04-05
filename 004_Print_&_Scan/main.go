package main

import "fmt"

func main() {
	/* Print */

	// Println() for a new line
	fmt.Println("We have met before, haven't we?")
	// Print() for printing in the same line
	// Use \n for line break
	fmt.Print("Next statement is in same line. ")
	fmt.Print("Ah shit...here we go again.\n")
	name := "Thomas Shelby"
	age := 34
	// Printf allows formatting
	// %v specifies a variable, a generic format specifier
	fmt.Printf("The name is %v. Age %v.\n", name, age)
	// %d for int, %s for string, %c for character, %f for float
	price := 45.2345
	// Float can be formatted to include number of digits after decimal
	fmt.Printf("The price is %0.2f\n", price)

	/* Sprint */

	// Sprint() saves a normal format string
	s1 := fmt.Sprint("The name is ", name, " and age is ", age)
	fmt.Println(s1)
	// Sprintf() saves a formatted string
	s2 := fmt.Sprintf("The name is %s and the age is %d.", name, age)
	fmt.Println(s2)
	// Sprintln() saves a string with new line
	s3 := fmt.Sprintln("Let's find out!")
	fmt.Print(s3)

	/* Scan */

	// Scan() stores the value entered via console
	var firstname string
	var lastname string
	fmt.Println("Enter firstname: ")
	fmt.Scan(&firstname)
	fmt.Println("Enter lastname: ")
	fmt.Scan(&lastname)
	fmt.Printf("The full name is %v %v.\n", firstname, lastname)

}
