# #10 Maps

Maps in Golang can be thought of to be very similar to Python dictionaries. They are of the key/value pair format.

## make() 

You can create a map using make(). Specify the type of key and type of value to do so.

```go
// create map with key of type string and value of type uint8
m1 := make(map[string]uint8)

// create map with key of type uint8 and value of type string
m2 := make(map[uint8]string)
```

## Accessing and assigning values

You can use `map[key]` syntax to access the value of a spefici key. The same syntax can also be used to assign a new value. If `key` does not exist, it will be made.

```go
m1["a"] = 4 // key a does not exist; make new and assign 4
fmt.Println(m["a"]) // prints 4
```

## len()

You can print length of map using the len() funtion.

```go
fmt.Println(len(m1))
```

## delete()

You can delete a specific key from the map using the delete() function. It takes two arguments: first is the map you want to delete from, second being the key you want to delete.

```go
delete(m1,"a")
```

## Check for a value

You can access a value, but also check if such a key exists in the map. That returns a boolean response.

```go
item,isPresent := m["c"]
if isPresent == true {
    fmt.Println(item)
} else {
    fmt.Println(isPresent)
}
```

## Alternate syntax

You can use the following syntax to declare and initialise a map.

```go
m2 := map[uint8]string{1:"Tom",2:"Jerry"}
fmt.Println(m2)
```

That's a wrap on maps. Day 10.