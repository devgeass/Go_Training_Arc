package main

import "fmt"

type employee struct {
	string
	float32
}

type student struct {
	id       uint
	name     string
	elective electiveSubject
}

type electiveSubject struct {
	subCode  uint
	subTitle string
}

type newStudent struct {
	id              uint
	name            string
	electiveSubject // note this!
}

func main() {
	// anonymous struct
	person1 := struct {
		firstname string
		lastname  string
	}{
		firstname: "Tony",
		lastname:  "Stark",
	}
	fmt.Println(person1.firstname + " " + person1.lastname)

	// anonymous fields
	employee1 := employee{
		string:  "John Doe",
		float32: 4500.25,
	}
	fmt.Printf("Name: %v | Salary: %v\n", employee1.string, employee1.float32)

	// nested structs
	student1 := student{
		id:   4,
		name: "Tim Cook",
		elective: electiveSubject{
			subCode:  244,
			subTitle: "Data Structures & Algorithms",
		},
	}
	fmt.Println(student1)
	fmt.Println(student1.elective.subCode)
	fmt.Println(student1.elective.subTitle)

	// promoted fields
	student2 := newStudent{
		id:   5,
		name: "Jack Hill",
		electiveSubject: electiveSubject{
			subCode:  345,
			subTitle: "Database Management System",
		},
	}
	fmt.Println(student2)
	fmt.Println(student2.subCode)  // promoted field!
	fmt.Println(student2.subTitle) // promoted field!
}
