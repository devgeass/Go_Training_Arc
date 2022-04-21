# #10 Range

The range function is used to iterate over data structures of Golang - array, slices or maps.

Syntax:

```go

for $1,$2 := range $3 {
    // code
}

where $1, $2 and $3 are arguments it takes. 

```

## Arrays and Slices

On arrays and slices you can iterate over each entry. The range function provides the index and item to iterate.

```go
cities := []string{"New York", "London", "Toronto", "Los Angeles"}

for index, city := range cities {
    fmt.Printf("%v: %v\n", index, city)
}
```

You can alternatively choose to skip the index. To do so, replace first argument with an underscore.

```go
for _, city := range cities {
    fmt.Println(city)
}
```

## Maps

In case of maps, range provides key and value pairs to iterate with.

```go
fruits := map[string]string{"a": "Apple", "b": "Banana", "c": "Cherry"}

for key, value := range fruits {
    fmt.Printf("%v -> %v\n", key, value)
}
```

You can alternatively choose to iterate over keys only which can be done by letting go of the second argument.

```go
for key := range fruits {
    fmt.Println(key)
}
```

That's a wrap on Day 11.