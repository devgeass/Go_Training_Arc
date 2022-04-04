# #3: Types

There are basic data types in Golang just like any other programming language.

## Unsigned Integer

- uint8 (or byte) range 0 to 255
- uint16 range 0 to 65535
- uint32 range 0 to 4294967295
- uint64 range 0 to 18446744073709551615
- uint (it will be 32 or 64 bit depending on the system)

## Signed Integer

- int8 range -128 to 127
- int16 range -32768 to 32767
- int32 (or rune) range -2147483648 to 2147483647
- int64 range -9223372036854775808 to 9223372036854775807
- int (it will be 32 or 64 bit depending on the system)

## Floating Number

- float32 
- float64

Note: There is no 'float' data type.

## Complex Number

- complex64
- complex128

Note: There is no 'complex' data type.

## String

Strings are a collection of bytes in Golang. For simplicity, a string is a collection of characters as we know it.

## Boolean

Boolean is a true/false value.

## Character or Rune

In Golang, a character is referred to as a rune and is determined as an int32 type. While a string is enclosed in double quotes, a character is enclosed in single quotes.


## Printing Type of a Variable

From Golang's *fmt* package, we get Printf() method (will cover different printing methods in upcoming chapter), wherein we can use %T format specifier to print type of a variable/constant. The '\n' signifies line break.

```go
var a = 10
fmt.Printf("%T\n",a)
```

## Type Conversion

If you want to perform operations between variables, they ought to be the **same type**. Not having so will result in error from Go regarding type mismatch.

```go
var i uint8 = 2
var j int32 = 45
var k int64 = 78
var l float32 = 90.25

fmt.Println(i+j) // error
fmt.Println(j+k) // error
fmt.Println(k+l) // error
```

To enable the operations, you must convert a variable to another type. If you convert the smaller type to a bigger type, it will be a lossless conversion. If you convert a bigger type to a smaller type, it will be a lossy conversion. The conversion function is the type itself.

```go
var i uint8 = 2
var j int32 = 45
var k int64 = 78
var l float32 = 90.25

fmt.Println(int32(i)+j) // 47
fmt.Println(int64(j)+k) // 123
fmt.Println(float32(k)+l) // 168.25
fmt.Println(k+int64(l)) // 168
```

That's a wrap on day 3. Keep learning!