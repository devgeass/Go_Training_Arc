# #8 Array

Array in Go is an ordered sequence of homogeneous items (having same data type) and having a fixed length. 

To declare and initialize an array,

```go
var_name := [length]type{values...}

data := [10]int{} // empty array
data := [2]int{2,3} // initialized with values
```
Note how length and datatype are compulsory in the array declaration syntax.

You can declare an array without initializing it. 

```go
var names [10]string // declaration
```

You can then use indexing to assign a value or update it. Index starts from 0. And indexing can be used to access value from the array.

```go
names[0] = "Tommy" // first element
names[0] = "Arthur" // update first element

fmt.Println(names[9]) // prints the last element
```

You can print the length of an array using the *len()* function.

```go
fmt.Println(len(names)) // prints 10
```

Note that len() prints the size of the array and not the number of elements present in it. 

If you declare an emtpy array, the default value of that data type will be assigned as all the elements.

```go
data := [3]int{}
fmt.Println(data)
// prints: [0 0 0]
```

Arrays can be multidimensional. You can declare 2 dimensional matrix as follows:

```go
matrix := [2][2]int8{{1,2},{3,4}} // 2x2 matrix
```

To access an element in 2D array, use indexing where first index represents a row and the second represents a column.

```go
fmt.Println(matrix[0][1]) // prints 2
```