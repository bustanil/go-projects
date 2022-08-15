package main

import (
	"fmt"
	"os"
)

func main() {
	deck, e := readFromFile("cards.txt")

	if e != nil {
		fmt.Println("Error: Failed to read file")
		os.Exit(1)
	}

	fmt.Println(deck)
	deck.shuffle()
	fmt.Println(deck)
}
