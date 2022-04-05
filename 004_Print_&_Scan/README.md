# #4: Print and Scan

While we have already seen a Println() function, Go has more to offer in the context of printing. Besides, we can also prompt users to enter values which we can store in variables to make our program dynamic.

Printing and scanning functions are offered by the *fmt* package.

## Printing

**Println()**

It is used to print to console and appends a new line after that.

```go
fmt.Println("Whatever...")
```

**Print()**

The print statement is simply used to print to console without a line break at the end.

```go
fmt.Print("Hey, I am in line 1.")
fmt.Print("Me too.")
```

**Printf()**

Uses format specifiers when printing a string, using which you can template a string with variables in it. There are different format specifiers depending on the data type. The most generic one is %v.

Here's a [cheatsheet](https://programming.guide/go/fmt-printf-reference-cheat-sheet.html) you can refer to.

```go
a := 2
b := 3
fmt.Printf("%v + %v = %v\n",a,b,a+b)
// prints 2 + 3 = 5 to the console 
```

**Sprint()**

It allows saving a print statement in the form of a string. You can add comma separated values to place variables.

```go
a := 2
b := 3
s := fmt.Sprint(a," + ", b, " = ", a+b)
fmt.Println(s)
// prints 2 + 3 = 5 to the console 
```

**Sprintf()**

Same as Sprint() but allows the use of format specifiers like Printf().

```go
a := 2
b := 3
s := fmt.Sprintf("%v + %v = %v",a,b,a+b)
fmt.Println(s)
// prints 2 + 3 = 5 to the console 
```

**Sprintln()**

Similar to Sprint() but it appends a new line.

## Scanning

Scan() function of Go can be used to take input via console during run time of the program. This helps create CLI applications.

You can declare a variable first. Then log a message using Println() and then use Scan() to read the data. The function accepts location of the variable which can be specified by prefixing the variable name using the & operator. `&var_name` points to the memory location of variable `var_name`. This concept is that of *pointers* and will be covered in future chapters.

```go
var firstname string
var lastname string
fmt.Println("Enter firstname: ")
fmt.Scan(&firstname)
fmt.Println("Enter lastname: ")
fmt.Scan(&lastname)
fmt.Printf("The full name is %v %v.\n", firstname, lastname)
```

That's a wrap on Day 4. Try practicing this module by creating a simple Q-A CLI application.