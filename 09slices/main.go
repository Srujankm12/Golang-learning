package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to the slices")

	// slice is a data structure that is a dynamic array

	var names = []string{"kd"}
	fmt.Printf("the type of a name is %T\n", names)

	names = append(names, "srujan", "jaii", "chethan", "nsk", "meghana")
	fmt.Println(names)

	var numbers = []int{10, 20, 30}
	fmt.Printf("the type of a number is %T\n", numbers)

	numbers = append(numbers, 200, 300)
	fmt.Println(numbers)

	numbers = append(numbers[1:4])
	fmt.Println("", numbers)

	highscores := make([]int, 4)
	highscores[0] = 100
	highscores[1] = 200
	highscores[2] = 300
	highscores[3] = 400

	highscores = append(highscores, 200, 300, 800)
	fmt.Println("", highscores)
	fmt.Println(sort.IntsAreSorted(highscores))
	sort.Ints(highscores)
	fmt.Println("", highscores)

}
