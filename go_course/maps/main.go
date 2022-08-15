package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors["red"])

	anotherColors := make(map[string]string)

	fmt.Println(anotherColors)

	colors["white"] = "#ffffff"
	fmt.Println(colors)

	delete(colors, "red")
	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
