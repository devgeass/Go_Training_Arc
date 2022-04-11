# #9 Slices

Slices don't provide constraint on number of elements like in the case of arrays. However, the type restriction still persists.

## make() function

The make() function can be used to declare a new slice with a desired length. The syntax is very similar to that of an array except you don't specify the fixed length inside [] like you used to with arrays.

```go
// make a slice of type uint8 having length 3
slice1 := make([]uint8,3)
```

However, this doesn't mean the length of `slice1` is restricted to size of 3. 

To add elements to the list, you can use indexing:

```go
slice1[0] = 0
slice1[1] = 1
```

## append() function

You can make use of the append() function to add elements to a slice.

```go
slice1 = append(slice1,2)
```

You can add multiple elements at the same time as well.

```go
slice1 = append(slice1,3,4,5)
```

## len() function

You can print the length of a slice using the len() function.

```go
fmt.Println(len(slice1))
```

## *slice* operator

There is an operator called slice operator which allows you retrieve a range of elements from a slice based on indexes. The syntax is of the form `slice[low:high]` which retrieves element starting from index `low` to index `high-1`. Index starts from 0, as you know.

```go
// assume slice1 = [0,1,2,3,4,5]
fmt.Println(slice1[1:4])
// prints: [1,2,3]
```

You can optionally skip an index.

```go
// assume slice1 = [0,1,2,3,4,5]
// print all elements after index 3
fmt.Println(slice[3:])
// prints: [3,4,5]
// print all elements before index 4
fmt.Println(slice[:4])
// prints: [0,1,2,3]
```

## Other way to declare and initialize

You can use alternative syntax to declare and initialize a slice which is very similar to that of an array.

```go
slice2 := []string{"a","b","c"}
```

## Multidimensional slice

Just like multidimensional arrays, multidimensional slices are also possible. However, they are more flexible given the fact that inner slices don't have to have the same length.

```go
slice3 := [][]uint8{{1},{2,3},{3,4,5}}
```

That's a wrap on day 9.