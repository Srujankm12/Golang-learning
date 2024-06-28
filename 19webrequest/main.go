package main

import (
	"fmt"
	"io"
	"net/http"
)

const URL = "https://www.google.com"

func main() {
	fmt.Println("HEY")

	response, err := http.Get(URL)
	if err != nil {
		panic(err)
	}
	fmt.Println("", response)
	response.Body.Close()
	// Ensure the response body is closed
	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(data)
	fmt.Println("", content)

}
