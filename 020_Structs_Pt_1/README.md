# #20 Structs Part 1

A struct is a user-defined type that represents a collection of fields.

## Declaration

Syntax:
```go
type struct_name struct {
    field_name type
    ...
}
```

Example:
```go
type user struct {
    firstname string,
    lastname  string,
    email     string,
    age       int
}
```

## Initialization

You can initialize a struct to a variable as you would any other data type. You have to provide values for the fields that the struct takes.

```go
user1 := user{
    firstname: "Thomas",
    lastname:  "Shelby",
    email:     "tshelby@gmail.com",
    age:       31,
}
```

You can alternatively skip writing the key names as long you enter everything in the right order.

```go
user1 := user{"Thomas","Shelby","tshelby@gmail.com",31}
```

It is okay to skip over fields when initializing a struct. You can do it later using the `dot` operator which is also used for access.

```go
user1 := user{"Thomas","Shelby","tshelby@gmail.com"}
user1.age = 31
```

## Access

You can access any individual field of struct using the `dot` operator.

```go
fmt.Println("User1 Firstname =", user1.firstname)
fmt.Println("User1 Lastname =", user1.lastname)
fmt.Println("User1 Email =", user1.email)
fmt.Println("User1 Age =", user1.age)
```

## Pointers at play

Suppose there is this function `createUser` that returns a pointer to a new struct.

```go
func createUser(firstname, lastname, email string, age uint16) *user {
	u := user{
		firstname: firstname,
		lastname:  lastname,
		email:     email,
		age:       age,
	}
	return &u
}
```

A new struct can then be created using the function.

```go
user2 := createUser("Arthur", "Shelby", "ashelby@gmail.com", 35)
```

In this case, `user2` is a pointer variable.

```go
fmt.Println(user2) // prints pointer to user2
fmt.Println(*user2) // prints value of user2
```

Alternatively you can declare a struct as a pointer using the `&` operator when initializing it.

```go
user3 := &user{firstname: "Ada", lastname: "Shelby", age: 28}
```

## Advantages of pointer

Since pointers deal with referencing, they can be used to alter value stored in a struct. Moreover, you can easily write modular code by writing a function that accepts pointer to the struct and make changes to it.

Consider the following `changeEmail` function. It takes two parameters: a pointer to user struct and new email address.

```go
func changeEmail(u *user, newEmail string) {
	u.email = newEmail
}

// passing pointer to user1 struct
changeEmail(&user1, "thomasshelby@gmail.com")
fmt.Println("User1 Email =", user1.email)

// user2 is already a pointer variable, remember?
changeEmail(user2, "arthurshelby@gmail.com")
fmt.Println("User2 Email =", user2.email)
```

That's a wrap on day 20.