package main

import "fmt"

func main() {
	fmt.Println("Lets learn pointers today")
	//The actual value of a pointer
	var ptr *int
	fmt.Println("", ptr)

	//The pointer variable
	newnumber := 21
	var ptr1 = &newnumber

	//the adresss of the pointer variable
	fmt.Println("", ptr1)

	//The value of the pointer variable
	fmt.Println("", *ptr1)

	//Adding the value

	*ptr1 = *ptr1 + 20
	fmt.Println("", *ptr1)

}
