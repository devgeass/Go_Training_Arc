# #7 Looping

Looping is used to repeat a certain task for finite or infinite number of times.

Go only has *for* loop. However, depending on the loop's construct, it's behavior changes.

## Conventional For Loop

The syntax for a finite for loop is no different than that of other programming languages.

```go
for initialize;condition;update {
    // code to loop
}
for i := 0; i <= 9; i++ {
	fmt.Printf("%v ", i)
}
```

## Condition Only For Loop

Sounds very similar to while loop, isn't it?

```go
for condition {
    // code to loop
}
```

## Infinite For Loop

Just neglect the condition and wrap your code in *for* block. That becomes an infinite loop.

```go
for {
    // code to infinite loop
}
```

## Continue Keyword

It is used inside a loop to suggest the loop to neglect the following expressions and continue with its next iteration.

```go
for j := 1; j <= 10; j++ {
	if j%2 == 0 {
		continue // skip below code if j is divisible by 2
	}
	fmt.Printf("%v ", j)
}
```

## Break Keyword

It is used to break out of the running loop. It can be used when some condition is met. It is used to break out of an infinite loop.

```go
for j := 1; j <= 10; j++ {
	if j == 3 {
		break
	}
	fmt.Printf("%v ", j)
}
```

That's a wrap on day 7.