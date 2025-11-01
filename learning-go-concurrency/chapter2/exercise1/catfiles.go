package main

import (
	"fmt"
	"os"
)

func main() {
	for i, a := range os.Args {
		if i == 0 {
			continue
		}

		content, err := os.ReadFile(a)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		fmt.Printf("File content of %s %s:", a, string(content))
	}
}
