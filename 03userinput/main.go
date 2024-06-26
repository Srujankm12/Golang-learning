package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	Welcome := "Welcome to Go Programming Language"
	fmt.Println(Welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")

	inputname, _ := reader.ReadString('\n') //error,ok
	fmt.Println("your name is:", inputname)

}
