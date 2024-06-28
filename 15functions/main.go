package main

import "fmt"

func main() {

	fmt.Println("welcome to function concept")

	srujan()
	Myname()
	result := adder(10, 20)
	fmt.Println("the result is ", result)
}

func Myname() {
	fmt.Println("my name is srujan")

}

func adder(valone int, valtwo int) int {
	return valone + valtwo
}

func srujan() {
	fmt.Println("i am srujan")
}
