package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Hello ")

	content := "hello too hey bee"

	file, err := os.Create("./test.txt")
	if err != nil {
		panic(err)
	}
	length, err := file.WriteString(content)
	if err != nil {
		panic(err)
	}
	fmt.Println("", length)
	file.Close()
}
