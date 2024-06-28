package main

import "fmt"

type user struct {
	name   string
	age    int
	email  string
	status bool
}

func main() {
	fmt.Println("STRUCT")

	Srujan := user{"srujan", 18, "srujankm12@gmail.com", false}
	fmt.Println(Srujan)
	fmt.Printf("complete details are :%+v\n", Srujan)
	Srujan.Getstatus()
	Srujan.getmail()
	Srujan.getage()
	Srujan.Getname()

}

func (u user) Getname() {
	fmt.Println("IS name is ", u.name)

}

func (u user) getmail() {
	fmt.Println("the email is", u.email)
}
func (u user) getage() {

	fmt.Println("the age is", u.age)
}

func (u user) Getstatus() {
	fmt.Println("the status", u.status)
}
