package main

import "fmt"

type user struct {
	firstname string
	lastname  string
	email     string
	age       uint16
}

func createUser(firstname, lastname, email string, age uint16) *user {
	u := user{
		firstname: firstname,
		lastname:  lastname,
		email:     email,
		age:       age,
	}
	return &u
}

func changeEmail(u *user, newEmail string) {
	u.email = newEmail
}

func main() {
	user1 := user{
		firstname: "Thomas",
		lastname:  "Shelby",
		email:     "tshelby@gmail.com",
		age:       31,
	}
	// user1 := user{"Thomas","Shelby","tshelby@gmail.com",31}

	fmt.Println("User1 =", user1)
	fmt.Println("User1 Firstname =", user1.firstname)
	fmt.Println("User1 Lastname =", user1.lastname)
	fmt.Println("User1 Email =", user1.email)
	fmt.Println("User1 Age =", user1.age)
	fmt.Println("- - Changing User1 Email - -")
	changeEmail(&user1, "thomasshelby@gmail.com")
	fmt.Println("User1 Email =", user1.email)

	user2 := createUser("Arthur", "Shelby", "ashelby@gmail.com", 35)
	fmt.Println("User2 =", user2)
	fmt.Println("*User2 =", *user2)
	fmt.Println("User2 Email =", user2.email)
	fmt.Println("- - Changing User2 Email - -")
	changeEmail(user2, "arthurshelby@gmail.com")
	fmt.Println("User2 Email =", user2.email)

	user3 := &user{firstname: "Ada", lastname: "Shelby", age: 28}
	fmt.Println("User3 =", user3)
	fmt.Println("- - Changing User3 Email - -")
	changeEmail(user3, "adashelby@gmail.com")
	fmt.Println("*User3 =", *user3)
}
