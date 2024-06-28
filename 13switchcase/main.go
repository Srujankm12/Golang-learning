package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator
	dicenumber := rand.Intn(6) + 1   // Generate a random number between 1 and 6
	fmt.Println(dicenumber)

	// Print the random number

	switch dicenumber {
	case 1:
		fmt.Println("You rolled a 1")
	case 2:
		fmt.Println("You rolled a 2")
	case 3:
		fmt.Println("You rolled a 3")
		fallthrough
	case 4:
		fmt.Println("You rolled a 4")
		fallthrough // fallthrough is used to execute the next case statement
	case 5:
		fmt.Println("You rolled a 5")
	case 6:
		fmt.Println("You rolled a 6")
	default:
		fmt.Println("You rolled a 7")
	}
}
