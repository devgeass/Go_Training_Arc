# #21 Structs Part 2

## Anonymous structs

Previously we saw that we declare structs using the `type name struct{}` syntax and then use it in the code. However, it is possible to create anonymous structs, i.e., structs with no name. These are used in cases like you only have to use this struct once.

To do so, use the following syntax:

```go
varname := struct {
    field type
    ...
}{
    corressponding_field value
    ...
}
```
You start by declaring the struct inline and then open another set of curly brackets and assign values to it. Here's an example.

```go
person1 := struct {
    firstname string
    lastname  string
}{
    firstname: "Tony",
    lastname:  "Stark",
}
fmt.Println(person1.firstname + " " + person1.lastname)
```

## Anonymous fields

It is possible to create structs with fields that contain only a type without the field name. These kinds of fields are called anonymous fields.

```go
type name struct{
    type
    type
    ...
}
```

Let's understand that with an example.

```go
type employee struct {
	string
	float32
}

employee1 := employee{
    string:  "John Doe",
    float32: 4500.25,
}
fmt.Printf("Name: %v | Salary: %v\n", employee1.string, employee1.float32)
```

The `employee` struct is declared with anonymous fields because the field names are missing. Only types are declared.

The variable `employee1` is of the type `employee` and to assigned the values to it, the nameless fields are used which is nothing but the types itself.

That is also the way to access a struct with anonymous fields.


## Nested structs

It is possible to have a struct reference another struct as its field.

Consider the following:

```go
type student struct {
	id       uint
	name     string
	elective electiveSubject
}

type electiveSubject struct {
	subCode  uint
	subTitle string
}
```

The struct `student` is referencing another struct `electiveSubject` as one of its fields `elective`.

Since a `struct` is a type in Golang, it can be used to assign a field.

Initializing nested structs is very similar to declaring non-nested structs.

```go
student1 := student{
    id:   4,
    name: "Tim Cook",
    elective: electiveSubject{
        subCode:  244,
        subTitle: "Data Structures & Algorithms",
    },
}
fmt.Println(student1.elective.subCode)
fmt.Println(student1.elective.subTitle)
```

## Promoted fields

Promoted fields is an advantage you get when you combine anonymous fields and nested structs.

Once again, anonymous fields are fields that are declared with type only, without name. Structs are type.

```go
type newStudent struct {
	id              uint
	name            string
	electiveSubject // note this!
}

type electiveSubject struct {
	subCode  uint
	subTitle string
}
```

The `newStudent` struct is same as `student` struct except for the fact that in `student` struct there was a field named `elective` whose type was `electiveSubject` whereas in `newStudent`, there is only the type `electiveSubject` without a name thereby making it an anonymous field.

The initialization is same as any other struct.

```go
student2 := newStudent{
    id:   5,
    name: "Jack Hill",
    electiveSubject: electiveSubject{
        subCode:  345,
        subTitle: "Database Management System",
    },
}
```

**But, what is the advantage you get by declaring a nested struct as an anonymous field?**

You get promoted fields.

Remember in the previous example of nested structs, to access the student's elective subject, we had to do `student.elective.subTitle`. This obvious because we are referencing the `electiveSubject` struct inside `student` struct and hence two `dot` operators are required to access inner properties.

But when you get promoted fields, by using anonymous struct field like in `newStudent`, you can skip double referencing with two `dot` operators like so:

```go
fmt.Println(student2.subCode)  // promoted field!
fmt.Println(student2.subTitle) // promoted field!
```

Notice we don't need to do `student2.electiveSubject.subTitle`, instead we can directly access properties of `electiveSubject` in `student` struct itself.

That's a wrap on day 21.