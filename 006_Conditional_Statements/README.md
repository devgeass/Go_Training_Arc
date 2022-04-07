# #6 Conditional Statements

Go has support for if-else statements and switch case.

## If-Else block

Traditional if-else block. No parenthesis are required when writing the condition.

```go
if condition {
    // code if condition is true
} else {
    // code if condition is false
}
```

You can have an if block without the else block. Else block is optional.

You can nest multiple conditions in if-else if ladder

```go
if condition {
    // ...
} else if condition {
    // ...
} else if condition {
    // ...
} else {
    // ...
}
```

Go allows a special syntax to declare variables preceding the condition which are available within the scope of the conditional block. Separate the expression and condition by a semicolon.

```go
if statement;condition {
    // ...
} else {
    // ...
}

//  statement is invalid outside the block
```

## Switch Case

Traditional switch case like any other programming language. However, you don't need to write the `break` keyword at the end of every case in Go.

You can have single case switch.

```go
switch condition {
    case case1:
        // code to execute for case1
    case case2:
        // code to execute for case2
    default:
        // code to execute if none of the cases match
}
```

Or you can have multi-case switch.

```go
switch condition {
    case case1,case2:
        // code to execute for either case1 or case2
    case case3,case4:
        // code to execute for either case3 or case4
    default:
        // default case
}
```

That's a wrap on day 6.