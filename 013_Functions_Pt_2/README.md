# #13 Functions Part 2

Functions can return multiple values or can be called with trailing arguments.

## Returning Multiple Values

As mentioned in previous part, a function returning a value must specify its return type. If you want to return multiple values, simply specify their return types enclosed in parenthesis. The types must in order as the returning values.

```go
func name(args) (type1, type2, type3) {
    return value1, value2, value3
}

// Here's an example
func add_sub(a, b int16) (int16, int16) {
	return a + b, a - b
}
```

Invoking the `add_sub` function above returns two values. These can be caught in variables in the following manner.

```go
sum, difference := add_sub(10, 20)
fmt.Printf("Sum: %v\nDifference: %v\n", sum, difference)
```

Futhermore, you can choose to ignore a particular returned value by replacing variable name with an underscore. The order is important.

```go
_, diff := add_sub(20, 10)
fmt.Printf("Diff: %v\n", diff)
```

## Variadic Functions

This is the name giving to functions which take trailing arguments, i.e., arbitrary number of arguments. However, they must be of the same type.

The `...` syntax is used to specify trailing arguments.

```go
func summation(numbers ...int) {
	fmt.Printf("%v -> sum: ", numbers)
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Printf("%v\n", sum)
}
```

Here's a function `summation` that takes arbitrary number of integer type arguments and returns its sum.

You can invoke the function with integers.

```go
summation(1, 4)
summation(1, 6, 10, 30)
```

You can otherwise sent a slice of integers to a variadic function as well. To do so, you have to use the `func(slice...)` syntax.

```go
arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
summation(arr...)
```

That's a wrap on day 13.