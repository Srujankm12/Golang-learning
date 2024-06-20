// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// )

//	func main() {
//		log.Println("Server is running at port 8000")
//		http.HandleFunc("/" , func(w http.ResponseWriter , r *http.Request){
//			fmt.Println("Server No Yes")
//		})
//		log.Fatal(http.ListenAndServe(":8000", nil))
//	}
package main

import "fmt"

func main() {
	var a string = "hello"
	Uppercase(&a)
	fmt.Println(a)
}

func Uppercase(s *string) {
	byt := []byte(*s)
	fmt.Println(byt)
	for i := range byt {
		if byt[i] >= 97 {
			byt[i] = byt[i] - 32
		}
	}
	*s = string(byt)

}
