# #2: Variables & Constants Part 2

There are some regulations you must understand when using variables and constants.

## Unused Variables

Go does not allow declaring variables and not using them **within a local scope**.

```go
func main(){
    var m1 int // error because it isn't used later
    m2 := 40 // error because it isn't used later
}
```

However, if you declare a variable in **global scope**, **it is fine if it remains unused**. Having said that, it is not a good idea at all to have unused variables/constants hanging around for no reason.

```go
package main

var d1 int // declared, not initialized, unused
var d2 = 50 // declared, initialized, unused

func main() { }
```

## Unused Constants

No problem at all. Unused constants throw no error when they remain unused, regardless of where they are declared or initialized.

```go
package main

const f1 = 60 // no error

func main() {
    const f2 int = 70 // no error
}
```

## Scope

A variable/constant declared within a block (a function, a conditional block  or a looping block) belongs to, and is accessible inside, that block and therefore it is locally scoped.

A variable/constant declared outside any blocks, at the package level, is globally scoped. Global variables could be public or private, meaning, they could be accessible outside the package or within the package respectively. This shall be covered when discussing packages.

If a variable/constant is declared in global scope, and a variable with the same name is declared in local scope, the preference is always given to local one.

```go
package main
import "fmt"
var d2 int = 50
const f1 = 50
func main() { 
    d2 := 60 // declared and initialized with shorthand
    const f1 = 80
    fmt.Println(f1)
    fmt.Println(d2)
}

$ go run main.go
80
60
```
The output of the above program is 80 and 60, because local scope is preferred over global scope.

## Multiple Declaration

You can declare multiple variables/constants at the same time. It can be done in three ways.

```go
var h1, h2 = 1, "Wow"
const PI, PREFIX = 3.14, "USER_"
```
You can declare variables or constants by separating them with a comma. The type of the variable is inferred by Go if you do not specify it.

You can declare multiple values by specifying a type. In doing so, you limit all the variables/constants to that particular type only.

```go
var k1, k2 int = 2, 3 // only int can be declared
```

Lastly, you can declare variables/constants by wrapping them inside parenthesis.
```go
var (
    c1 = 20 // initialized, type inferred
    c2 int = 30 // initialized, type specified
    c3 int // declared
    c4 = "Really?" // initialized, type inferred
)
c3 = 40
```
The same syntax goes for constants by replacing *var* with *const*.

## Naming Rules

- An identifier must not be a reserved keyword of Go.
- An identifier must start with a letter or an underscore.
- An identifier may not contain any special character other than underscore.
- An identifier beginning with a capital letter is exported (more about it later).

## Naming Conventions

-  Use short and precise names.
-  Stick with a particular naming case - snake case or camel case.
- Constants should be capitalized, separated by underscore if it is more than a word. 
  


That's a wrap on Day 2.