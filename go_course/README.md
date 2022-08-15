# Go: The Complete Developer’s Guide

This is my personal notes when learning Go from this [Udemy course](https://www.udemy.com/course/go-the-complete-developers-guide/learn/lecture/7824514?start=570#questions).

The sample code is located in a sub directory for each module.

# Package

package == project == workspace ??

package **main** is very special, it tells the compiler to create an executable for that .go file.

# Import

import is used to access other packages inside the source file you are writing.

standard packages can be found here [https://pkg.go.dev/std](https://pkg.go.dev/std)

# Variable Declaration

## var

```go
var name string = "Bustanil"
```

## :=

```go
name := "Bustanil"
```

# Function

## Function Definition

```go
func newCard() string {
	return "Five of Diamonds"
}
```

## Receiver Function

In Go, we can define a Receiver Function. A receiver function can only be called by the “receiver” type.

```go
// define a function with the receiver type 'deck'
func (d deck) print() {
	fmt.Println(d.toString())
}
```

## Multiple Return Values

On of the cool things in Go is that we can define a function that returns multiple values.

```go
func myFunction(fullName string) (string, string) {
	names := strings.Split(fullName, " ")
	return names[0], names[1]
}
```

# Array and Slice

## Array

fixed length list

## Slice

An array that can grow or shrink

### Creating a new slice

```go
cards := []string{"Ace of Diamonds", "Six of Spades"}
```

### Appending a slice

```go
cards = append(cards, "Six of Hearts")
```

**Note:** remember to assign the return value of the `append` function.

### Iterating over a slice

```go
for i, card := range cards {
	fmt.Println(i, card)
}

// we can ignore one of the return value by using _
// example: ignore the index i
for _, card := range cards {
	fmt.Println(card)
}
```

### Sub Slice

We can get a sub slice from a slice using the following syntax

```go
mySubSlice := mySlice[start:end]
// start inclusive, end is exclusive (not included)

// we can omit the either the start
mySubSlice := mySlice[start:] // create a sub slice from start to end of the slice

// we can also omit the start
mySubSlice := mySlice[:end] // create a sub slice from 0 to end
```

Example:

```go
fruits := []string {"Apple", "Banana", "Mango", "Durian"}
myFruits := fruits[0:2] // "Apple", "Banana"
```

# Type

## Type Alias

We can create a type alias

```go
// create a new type called 'deck' which is a Slice of string
type deck []string
```

## Type Conversion

We can convert a value of one type to another type by calling the target type as if it’s a function.

Example:

```go
var myByteSlice []byte = []byte("Hello world")
```

# IO

## ioutil

### WriteFile

```go
ioutil.WriteFile("myfile.txt", []byte("Bustanil"), 0666)
```

### ReadFile

```go
var bs []byte, err := ioutil.ReadFile("myfile.txt")
```

# Random

## Generating Random Number

1. Create a Source

```go
source := rand.NewSource(time.Now().UnixNano())
```

1. Create a Rand from the Source

```go
r := rand.New(source)
```

1. Generate the random number

```go
r.Intn(len(d) - 1)
```

# Testing

Files ended with `_test.go` are considered as test files. To run the tests for a particular package we can use the following command:

```go
go test
```

Inside the test file, there should be at least one function that starts with `Test` and it should receive a parameter of type `testing.T`. Example:

```go
func TestNewDeck(t *testing.T) {
	...
}
```

When writing the test, we surely want to assert one or more values, there’s no special method for that. We can just use an `if` to test the value and report an error if it’s false.

```go
if d[0] != "Ace of Spades" {
	t.Errorf("Unexpected value", d[0])
}
```

# Struct

We can define a custom composite type using `struct`

## Declaring a struct

```go
type person struct {
	firstName string
	lastName string
}
```

## Instantiating a struct

### Option #1

```go
jim := person{"Jim", "Carey"}
// NOTE: the order matters
```

### Option #2

```go
jim := person{firstName: "Jim", lastName: "Carey"}

// order does not matter as we specify the field names
jim = person{lastName: "Carey", firstName: "Jim"}
```

## Updating a struct

```go
jim.firstName = "Jimmy"
```

## Print struct value with their field names

```go
fmt.Printf("%+v", jim)
// {firstName:Jim lastName:Carey}
```

## Embedding struct

We can embed a struct inside another struct

```go
type struct contactInfo {
	email string
	zipCode int
}

type struct person {
	firstName string
	lastName string
	contact contactInfo
}
```

We can omit the field name, and the field name will follow the type name

```go
type struct person {
	firstName string
	lastName string
	contactInfo // equivalent to: contactInfo contactInfo
}
```

## Pass by value

When used as a function receiver, the struct that is available to the function is a **copy** of the receiver. That means any modification to the struct is only effective inside the function, the original instance is left unchanged.

# Pointer 😥

By default, when we pass a value of `int`, `float64`, `bool`, `string` and `struct` into a function either as receiver or as parameter, the value will be passed by value (copy).

Especially when dealing with struct if we want to modify the original value (not the copy) the we have to use a pointer.

So for example, if we want our function to be able to modify the receiver struct then we have to use the pointer operator.

```go
func (p *person) updateFirstName(newFirstName string) {
	(*p).firstName = newFirstName;
}
```

Notice that we use `(*p)` to dereference the pointer.

We can then use the function to update the fields inside a `person` struct.

```go
p := person{firstName: "Alex", lastName: "Anderson"}
p.updateFirstName("John")

fmt.Println(p) // {"John", "Anderson"}
```

## Pass by Reference

The other types such as `slice`, `map`, `channel`, `pointer` and `function` are by default pass by reference. Meaning that we don’t need to use pointers to be able to modify value of these types.

```go
type deck []string

// here we do not have to use pointer operator as deck is an alias of []string
// and []string is always pass by reference
func (d deck) setFirstCard(value string) {
	d[0] = value
}
```

# Map

Map is a key-value data structure.

## Declaring a Map

```go
var myMap := map[string]int{
	"red" : 1,
	"blue": 2,
	"green": 3,
}
```

## Modifying Map

### Updating the value

```go
myMap["red"] = 10
```

### Deleting a key

```go
delete(myMap, "red")
```

## Iterating over Map

```go
for color, value := range myMap {
	fmt.Println(color, value)
}
```

# Interface

An interface defines a contract of how one or more behaviors (functions) should be (signature).

```go
type bot interface {
	getGreeting() string
}
```

To “implement” the interface we just need to define a receiver function which matches the `getGreeting()` signature.

```go
type englishBot struct{}

func (englishBot) getGreeting() string {
	return "Hello friend!"
}
```

Now we can use the interface

```go
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	eb := englishBot{}

	// here we can pass 'eb' because there's a receiver function
	// for type 'englishBot' which conforms the 'bot' interface
	printGreeting(eb) // prints "Hello friend!"
}
```

**Gotchas:**

Defining a receiver function for `englishBot` is different from `*englishBot`

# Goroutine

Goroutine is the main mechanism in Go to achieve concurrency.

```go
func doSomething() {
	fmt.Println("Hello")

	go func() {
		fmt.Println("world") // will be run in background
	}()

	// do another thing 
	fmt.Println("John")
	time.Sleep(1 * time.Second) // give time for the goroutine to run
}
```

```go
// output: 
Hello 
John 
world
```

Go Routine is a lightweight thread. It’s cheaper than the real OS thread.

Goroutines communicate with each other using `channel`

To create a channel we have to use the built-in `make` function.

```go
c := make(chan string) // create a string channel
```

To write to a channel use the operator `channel<-`

To read from a channel use the operator `<-channel`()

Example use:

```go
func doSomething() {
	fmt.Println("Hello")
	c := make(chan string)

	go func() {
		c<- "world" // write to the channel
	}()
	
	fmt.Println(<-c) // read from channel, blocks until there's something to read
}
```

```go
// output:
Hello
world
```