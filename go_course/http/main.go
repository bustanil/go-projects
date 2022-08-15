package main

import (
	"net/http"
	"os"
	"fmt"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Printf("%+v", resp)

	buffer := make([]byte, 64)
	_, err = resp.Body.Read(buffer)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println(string(buffer))
	resp.Body.Close()
}
