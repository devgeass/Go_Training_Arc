# #12 Functions

Functions help write modular code. They are used to split logic into smaller and manageable code.

Syntax:

```go
func function_name(args) return_type {
    // code
}
```

The _func_ keyword is used to declare a function. It may or may not take any arguments. It may or may not return any value.

## No Argument No Return

```go
func greet() {
	fmt.Println("Have a nice day!")
}
```

## With Argument No Return

```go
func greetByName(name string) {
	fmt.Printf("Have a nice day, %v!\n", name)
}
```
You must specify the type of the argument that you pass.

## No Argument With Return

```go
func nani() string {
	return "Omae wa mou shindeiru!"
}
```
You must explicitely specify the return type.

## With Argument With Return

```go
func add(a, b uint8) uint8 {
	return a + b
}
```
If multiple arguments are of the same type, you can declare the type on final parameter.

## Invoking functions

You can invoke functions in Go using the following syntax:

```go
function_name(args)
```

If a function returns a value, you can store it in a variable.

That's a wrap on day 12.