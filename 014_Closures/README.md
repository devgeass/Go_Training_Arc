# #14 Closures

A closure is the combination of a function bundled together (enclosed) with references to its surrounding state. In other words, a closure gives you access to an outer function's scope from an inner function.


Go allows creating anonymous functions, i.e., inline nameless functions. This can be used when creating closures. Have a look at the example below.

## Understanding Closures

```go
func makeAdder(x int) func(int) int {
	return func(y int) int {
		return x + y
	}
}
```

The function `makeAdder` takes an integer argument and returns a function which takes an integer argument which in turn returns an integer.

```
makeAdder(x int) func(int) int

makeAdder       function name
(x int)         integer argument
func(int) int   return type
```

The anonymous function is returned inside `makeAdder` having the function definition matching the return type of parent function.

```
return func(y int) int

func(y int)     anonymous function
(y int)         integer argument
int             return type
```

The inner function can access variable in outer scope, the one belonging to its parent function. This is *lexical scoping*. Also, a thing to note is, the value of outer function's variables is stored in its own lexical environment. Let's understand that with the following example.

```go
add5 := makeAdder(5)
fmt.Println(add5(4))
fmt.Println(add5(10))
```

Here we create a closure function `add5` which passes the argument 5 to `makeAdder`. This creates a closure on `add5` whose lexical environment store `x=5` for its future reference.

Now, the inner function of the closure can be invoked with the closure function. We know that inner function takes an integer argument. Hence, `add5(4)` calls the inner anonymous function with `y=5` and since `x=5` was already stored in `add5`'s lexical environment, the `add5(4)` returns 9. 

Similarly, you can create another closure function based on the same `makeAdder` definition.

```go
add10 := makeAdder(10)
fmt.Println(add10(4))
fmt.Println(add10(10))
```

Now, `add10` is closure function with lexical environment storing `x=10`. There on any calls made to the inner function will always register the value of x to be 10.

As you notice, the `makeAdder` function now serves like a function factory that hands out a function with initial value to which you can perform add operation using another argument.


## Usefulness

Closures are useful because they let you associate data (the lexical environment) with a function that operates on that data.

Closures can be used when you want to create prototype functions like `makeAdder` in the above example based on which other functions can be derived. This behaviour is very similar to that of Object Oriented Programming where objects allow you to associate data (the object's properties) with one or more methods.

That's a wrap on day 14.