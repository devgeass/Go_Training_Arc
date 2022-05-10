# #18 BuiltIn Functions Part 4

## Parsing

There are times when you need to convert string type to other types. The `strconv` package comes in handy to help achieve this conversion.

| Function   | Syntax                  | Purpose                                                            |
| ---------- | ----------------------- | ------------------------------------------------------------------ |
| parseInt   | parseInt(s, base, bits) | convert string s to int with given base and no. of bits            |
| parseBool  | parseBool(s)            | converts values- t,1,T,true,True,TRUE,f,0,F,false,FALSE to boolean |
| parseFloat | parseFloat(s, bits)     | converts string s to float with given no. of bits                  |
| Atoi       | Atoi(s)                 | converts string s to integer base 10                               |
| Itoa       | Itoa(n)                 | converts int n to string                                           |

The `strconv` function returns two values. One is desired result and other is error value. Incase an error occurs, you can always deal with it in an appropriate way. The case can occur if invalid input has been passed.

```go
n, err := strconv.parseInt("50",10,64)
if err != nil {
    // error handling code
}
// ...
```

## Logging

The log package provides a logger which is helpful when tracking of statement execution is required. It prints date and time everytime it logs something to the console. It provides three main functions, each with three variations: 

1. Print[f|ln]
2. Fatal[f|ln]
3. Panic[f|ln]


The `log.Print()`, `log.Println()` and `log.Printf()` do the same thing as the ones provided by `fmt` except logger adds date and time to them.

```go
log.Println("Welcome to Go Learning Arc by Devgeass!")
```

The `log.Fatal()`, `log.Fatalln()` and `log.Fatalf()` are used when you want to throw an error message and exit the program. Fatal functions in turn call the `os.Exit(1)` to stop program execution after printing the error message.

```go
func divide(a, b float32) float32 {
	if b == 0 {
		log.Fatal("Cannot divde by zero!\n")
	}
	return a / b
}
```

The function `divide` will stop the program execution if `b` argument is passed value of zero.

The `log.Panic()`, `log.Panicln()` and `log.Panicf()` are used when exceptions can occur in the code during run time and program needs to be terminated. Panic functions in turn call the `panic()` function and terminate the program. More information on `panic()` will be covered in upcoming chapters.

That's a wrap on day 18.