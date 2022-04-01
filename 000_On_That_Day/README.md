# #0 On That Day

Let's GO! (no pun intended)

Download and install Go on your local machine by following one of the many amazing tutorials on YouTube. I am using VS Code as an editor (enabled Intellisense because why not?)

There is also [GO Playground](https://go.dev/play/) for browser based code development.

## Now then... first things first!

- Go is a compiled language.
- We get to use the terminal for executing commands.
- It is statically typed.
- It is a high performance language.
- One of the major highlights is *goroutines* - Go's support for concurrency and parallelism. 
- Docker and Kubernetes are built on Go.

## Sample Program

The complete sample program is in the folder. I assume you have prior knowledge of other programming languages like Java, C, C++, Python or JavaScript. Relate with it.

```go
package main
```
Every Go program starts with a *package declaration*. Packages help in organizing your code. In our code, we are declaring that *main.go* code belongs to the **main** package.    

```go
import "fmt"
```
The *import* statement is used to import packages in the current file. The **fmt** package is the *format* package which provides formatting functions (an obvious guess).

```go
func main(){

}
```
Well that's a function. Execution starts from *main* function.

```go
// Comment symbol is the same!
/* Even this one!
What a relief!
*/
```
Commenting is same old.

```go
fmt.Println("What's up, danger?")
```
From the *fmt* package, the *Println* function is used to print to the console. Please don't print "Hello World!" Also, we have to use double quotes " ". 

## Fun Program 

Run* the program I mean.

To run the program, execute from your terminal/CMD, within the current directory (where main.go file is):
```sh
go run main.go
``` 
And that should log "What's up, danger?" If not, well... we need to talk.

That's all there is to it. That's a wrap on Day 0.

## No! Wait!

Apparently, we can create an executable of the program. Isn't that convenient? To do so, execute in the working directory:
```sh
go build main.go
```
That should produce an executable in the parent directory.

Now that is a wrap on Day 0.
