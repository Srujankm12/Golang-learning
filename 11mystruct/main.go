package main

import "fmt"

type user struct {
	name  string
	age   int
	email string
}

func main() {
	fmt.Println("STRUCT")

	Srujan := user{
		"srujan", 18, "srujankm12@gmail.com",
	}
	fmt.Println(Srujan)
	fmt.Printf("complete details are :%+v\n", Srujan)

}
