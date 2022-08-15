package main

import "fmt"

type englishBot struct{}
type spanishBot struct{}

type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an english greeting
	return "Hi There!"
}

// we can omit the receiver name if we do not use it
func (spanishBot) getGreeting() string {
	// VERY custom logic for generating a spanish greeting
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
