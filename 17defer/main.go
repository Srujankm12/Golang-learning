package main

import "fmt"

func main() {

	defer fmt.Println("I am in the defer")
	defer fmt.Println("jaii")
	fmt.Println("Hello World")
	defer fmt.Println("ho")

	mydefer()
}
func mydefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println("", i)
	}
}
