package main

import (
	"fmt"
)

func main() {

	fmt.Println("loops in go")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	fmt.Println("", days)

	for d := 0; d < len(days); d++ {
		fmt.Println("day is ", days[d])
		fmt.Println("hey hey wait\n")
		for i := range days {
			fmt.Println("day is ", days[i])

			rougevalue := 1
			for rougevalue < 10 {
				fmt.Println("rougevalue is ", rougevalue)
				rougevalue++

			}
		}

	}
}
