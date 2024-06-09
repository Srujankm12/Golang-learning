package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to the time")

	presenttime := time.Now()
	fmt.Println(presenttime)

	fmt.Println(presenttime.Format("01-02-2006 Monday 15:04:05 "))

	createddate := time.Date(2021, time.April, 9, 20, 9, 2, 0, time.UTC)
	fmt.Println(createddate.Format("01-02-2006 Monday  "))

}
