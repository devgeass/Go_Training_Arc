# #17 BuiltIn Functions Part 3

## Array

You can find the length of the array using `len(array)` function.

```go
arr := [3]uint8{1,2,3}
fmt.Println(len(arr))
```

There is a shorthand syntax for array where you don't have to specify the size. However, once initialised with the elements, the size will not change.

```go
arr1 := [...]uint8{1,2,3,4,5}
arr1[6] = 6 // throws error because size is 5
```

You can copy an array to another variable simply by assigning it to new variable.

```go
arr2 := arr
```

Modifying the copy array does not affect the original array in Golang because unlike other languages where arrays are stored as references, in Golang they are stored as values.

```go
arr2[0] = 10
fmt.Println(arr2) // logs: [10 2 3]
fmt.Println(arr) // logs: [1 2 3]
```

## Slice

You can find the length of the slice using `len(slice)` function.

```go
slc := []uint8{1,2,3}
fmt.Println(len(slc))
```

Append new elements to a slice using `append()` function.

```go
slc = append(slc,4,5,6)
```

Append another slice to a slice using `append()` function but to do that you need the spread `...` operator.

```go
slc2 := []uint8{7,8}
slc = append(slc, slc2...) 
```

You can copy one slice to another variable by assigning it. However, any change in copy slice will affect original slice because slices are stored as references.

```go
slc3 := slc2
slc3[0] = 9
fmt.Println(slc3) // logs: [9 8]
fmt.Println(slc2) // logs: [9 8]
```

## Map

You can find the length of the map using `len(map)` function.

```go
map1 := map[string]uint8{"a":1,"b":2}
fmt.Println(len(map1))
```

You can copy one map to another variable by assigning it. However, any change in copy map will affect original map because maps are stored as references.

```go
map2 := map1
map2["a"] = 3
fmt.Println(map2) // logs: map[a:3 b:2]
fmt.Println(map1) // logs: map[a:3 b:2]
delete(map2,"a")
fmt.Println(map2) // logs: map[b:2]
fmt.Println(map1) // logs: map[b:2]
```

That's a wrap on day 17.