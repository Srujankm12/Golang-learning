package main

import "fmt"

func main() {
	fmt.Println("If else in go")

	logincount := 10
	var result string

	if logincount < 9 {
		result = "user is regular"

	} else {
		result = "user is notregular"
	}
	fmt.Println("", result)

	if 9%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
	if num := 20; num < 10 {
		fmt.Println("the number is less than 10")
	} else {
		fmt.Println("the number is greater than 10")
	}

}
