package main

import "fmt"

func main() {

	fmt.Println("Lets learn array today")

	var age [3]int

	age[0] = 20
	age[2] = 21
	age[1] = 22
	fmt.Println("", age)

	var names [5]string

	names[0] = "srujan"
	names[3] = "sachin"
	names[2] = "dhoni"
	names[4] = "hhh"
	fmt.Println("", len(names))

	var balls [3]int
	balls = [3]int{1, 2, 3}
	fmt.Println("", balls)
	fmt.Println("", len(balls))

}
