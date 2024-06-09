// package main

// import (
// 	//"bufio"
// 	"bufio"
// 	"fmt"
// 	"os"
// 	//"os"
// )

// func main() {

// 	fmt.Println("Welcome to my coding journey")
// 	var name string = "Srujan"
// 	fmt.Println("My name is ", name)
// 	fmt.Printf("type is %T\n", name)

// 	var age = 20
// 	fmt.Println("mys age is ", age)
// 	fmt.Printf("type is %T\n", age)

// 	// reader := bufio.NewReader(os.Stdin)
// 	// fmt.Println("Enter your name:")

// 	// inputname, _ := reader.ReadString('\n')
// 	// fmt.Printf("type is %T\n", inputname)
// 	// fmt.Println("name is ", inputname)

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter your name:")

// 	inputname, _ := reader.ReadString('\n')
// 	fmt.Println("Entered name is", inputname)
// 	fmt.Println(inputname)

// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("welcome to rating shop")
	fmt.Println("give rating from 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("your rating is ", input)

	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("error occured", err)
	} else {
		fmt.Println("your rating is ", numrating+1)
	}

}
