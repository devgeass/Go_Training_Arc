# #19 Pointers

Pointers are variables that store address (of another variable). Two symbols to understand absolutely when learning pointers are:

1. Ampersand(`&`): Known as referencing operator.
2. Asterisk(`*`): Known as deferencing operator.

## Declaration of Pointer

The declaration syntax uses asterisk(*) operator : `var var_name *datatype`

```go
var ptr *int
```

## Initializing a pointer

Note that pointer variables store address of other variable. To obtain the address of any variable we use the ampersand(&) operator, thereby getting the *reference* of that variable meaning its memory location. This can be then assigned to the pointer variable.

For example, if `a=5`, address of `a` can be obtained using `&a`. 

```go
var i int = 64
// pointer variable j stores address of i
var j *int = &i

// shorthand syntax works too
st := "devgeass"
str := &st
```

That's how the reference operator `&` is used.

## Referencing and Deferencing

Referencing refers to pointing to the memory location using the `&` operator. You can interpret `&var` as the question "What is the address of var"?

Deferencing is the opposite. Deferencing operator `*` is used on pointer variables to access the value they are pointing at. That means, it is used on any variable that has address stored or an address in general.

Syntax: `*(address_variable)` or `*(address)`

But what is an address variable? That's right. A pointer.

Therefore, if `a=5`, we know `&a` stores address of a. Hence, if you want to deference it, use `*` to obtain the value stored in address of a which is nothing by 5 using `*(&a)`. 

To put simply, the question to ask when using dereferencing operator on `*var` is "What is the content stored at address `var`?" knowing that `var` is either a pointer variable or an addresss.

```go
i := 10
j := &i
fmt.Println(*j) // logs 10
fmt.Println(*&i) // logs 10
```

## Example

```go
01 | var m int = 73
02 | // m := 73
03 | var n *int = &m
04 | // n := &m
05 | fmt.Println(m)
06 | fmt.Println(&m)
07 | fmt.Println(n)
08 | fmt.Println(*n)
09 | fmt.Println(&n)
10 | fmt.Println(*&m)
```

Lines 1 and 3 declare two variables - `m`, an integer and `n`, an integer pointer that stores location of `m`.

Lines 2 and 4 are alternate shorthand declaration syntax of lines 1 and 3 respectively.

Line 5 simply prints value of `m` to the console. Therefore, it prints 73.

Line 6 prints the reference of `m`, i.e., the location of `m` in the memory. Let's say for example that is `0xc40`.

Line 7 prints `n` to the console, which as we know is nothing but a pointer to location of `m` (line 3). Therefore, `n` will print location of `m` same as line 6, i.e., `0xc40`.

Line 8 uses deference operator on `n`, therefore, it prints the value stored in address stored in `n`. As stated above, the deference operator is used to obtain content stored at a specific address in the memory. In this case, content at `n` is content at address of `m` which is value of `m` which is 73. Therefore, line 8 prints 73 to the console.

Line 9 prints the reference of `n`, i.e., its memory location. Despite the fact that `n` is a pointer variable, in itself it's a variable that needs to be stored in the memory and therefore will have some memory location as well. In other words, pointers store address of another variable but also do have a particular address themselves. Hence, `&n` prints address of `n` in memory which let's say for example is `0xd67`.

Line 10 deferences the reference of `m`. The easiest way to remember this is that the deference operator and reference operator cancel each other, because they are opposite in nature, and therefore all we are left with is `m`. That's exactly what line 10 prints, the value of `m`, 73.

## Array Example

In previous modules, we studied that arrays in Go are stored as values and not references. This means, when we create a copy of the array in another variable, any change to the copy variable will not affect the original variable.

However, this behaviour can be influenced with the use of pointers.

```go
arr := [3]int{1, 2, 3}
// copy arr
arr_copy := arr
// modify copy arr
arr_copy[0] = 10
fmt.Println("arr_copy =", arr_copy)
// no change in original arr
fmt.Println("arr =", arr)
// copy arr to pointer variable
// var arr_copy_ptr *[3]int = &arr
arr_copy_ptr := &arr
// modify pointer type arr value
arr_copy_ptr[0] = 20
// print value stored in address stored by arr_copy_ptr -> arr
fmt.Println("arr_copy_ptr =", *arr_copy_ptr)
// print arr, it is modified
fmt.Println("arr =", arr)
```

Here the change is made directly to the memory location using pointer variable and therefore, the original array is also affected. This is how powerful pointers are.

## Pointer Chaining

As stated earlier, a pointer variable is also a variable which is stored somewhere in the memory, meaning it has an address. Therefore, it is possible to have a pointer to a pointer!

```go
i := "devgeass"
j := &i // pointer to string i
k := &j // pointer to pointer j
fmt.Println("i =", i) // value of i, devgeass
fmt.Println("&i =", &i) // address of i
fmt.Println("j =", j) // value of j is reference of i, therefore, address of i
fmt.Println("*j =", *j) // deference of j, therefore, value of i, devgeass
fmt.Println("&j =", &j) // address of j
fmt.Println("k =", k) // value of k is reference of j, therefore, address of j
fmt.Println("*k =", *k) // deference of k, therefore, value of j, therefore address of i
fmt.Println("**k =", **k) // deference of deference of k is value of i, devgeass
```

`**k` might seem a little confusing at first but if you have understood the concepts of pointers well, you should not have any problem in figuring out as to how `**k` results in value of `i`.

That's a wrap on day 19. Pointers are very crucial to memory management.