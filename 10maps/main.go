package main

import (
	"fmt"
)

func main() {

	fmt.Println("Welcome to maps")

	languages := make(map[string]string)

	languages["Go"] = "Golang"
	languages["Js"] = "Javascript"
	languages["Py"] = "Python"
	fmt.Println(languages)
	fmt.Println("", languages["Js"])
	fmt.Println("", languages["Go"])
	fmt.Println("", languages["Py"])

	delete(languages, "Go")
	fmt.Println("", languages)

	for key, value := range languages {
		fmt.Printf("for key %v ,value is %v\n", key, value)
	}
}
