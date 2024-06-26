package main

import "fmt"

//const LoginToken = "HEYHEYHEY" //publicc

func main() {

	var username = "SRUjan"

	fmt.Println("THe string is ", username)
	fmt.Printf("The variables of the type is: %T\n", username)

	var age = 90
	fmt.Println("the kiran age is", age)
	fmt.Printf("the variable is of type is :%T\n", age)

	var kk bool = false
	fmt.Println("the bool is ", kk)
	fmt.Printf("The variables of the type is : %T\n", kk)

	var smallvalue uint8 = 255
	fmt.Println("The age of the student is ", smallvalue)
	fmt.Printf("The variables of the type is:%T\n", smallvalue)

	var floatvalue float32 = 200.0909238983
	fmt.Println(floatvalue)
	fmt.Printf("The variables of the type is %T\n", floatvalue)

	var anothervalue int
	fmt.Println("the value is ", anothervalue)
	fmt.Printf("The variables of the type is %T\n", anothervalue)

	//
	numberofusers := 10000
	fmt.Println(numberofusers)
	fmt.Printf("The variables of the type is %T\n", numberofusers)

	//fmt.Println(LoginToken)
}
