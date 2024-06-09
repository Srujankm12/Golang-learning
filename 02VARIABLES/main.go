package main

import "fmt"

//const LoginToken = "HEYHEYHEY" //publicc

func main() {

	var username string = "Srujan"
	fmt.Println("The name of Student is ", username)
	fmt.Printf("The variables of the type is: %T\n", username)

	var age int = 21
	fmt.Println("The age of the student is ", age)
	fmt.Printf("The variables of the type is:%T\n", age)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("The variables of the type is : %T\n", isLoggedin)

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
