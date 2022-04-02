# #1 Variables & Constants Part 1

Like any other programming languages out there, variables and constants are, as we know, memory locations having some value of a specified type, identified with a name. Values in variables can be changed. Constants are well...constant.

## Variables

Declaring and initializing a variable in Go is done using the *var* keyword. The most basic syntax is : *var* variable_name = value 


```go
var v1 = 10
```
Go is a statically typed language. It means every variable or constant is associated with a specific type that cannot be changed after its declaration. 

If you do not specify a type manually, Go automatically infers the type.
```go
v1 = "Transform!" // throws error
v1 = 15 // valid
```

On the other hand, you can declare and initialize a variable while also specifying the type on your own. The syntax is : *var* variable_name type = value
```go
var v2 int = 20
``` 
You can declare a variable first and then assign it a value later. However, you must specify the type when doing so in order to achieve type safety.
```go
var v3 int
v3 = 30
v3 = "Transform!" // throws error for type mismatch
``` 
You can declare variables outside a function in a global scope. The syntax will not change...obviously.
```go
package main

var d1 int
var d2 = 50

func main() {
    d1 = 50
 }
```
You can definitely, declare variables in global scope but then assign value inside a function.

Do your love shortcuts? Go got you covered with shorthand syntax using := 
```go
v4 := 40
```
Using this syntax infers the type on its own which may be a good thing or a bad thing depending on the use case. GO has multiple types (will be covered in the  upcoming chapter) and therefore, to use a custom type you cannot use this syntax.

**Note: The shorthand syntax cannot be used in global scope.**

```go
package main 

v4 := 40 // throws error

func main() { }
```

## Constants

Constants are declared with the *const* keyword with the syntax : *const* variable_name = value
```go
const f1 = 70
```
The type is inferred by Go, the same case as with variables. And similarly, you can declare constants while specifying the type.
```go
const f1 int = 70
```
When declaring a constant it is mandatory to provide a value, which then cannot be changed because obviously, once again, it's called a constant for a reason.

You can declare constants in function scope or in global scope.

That's a wrap on Day 1. There's a part two for Variables & Constants chapter to discuss a few more things.