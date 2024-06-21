package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

	})

	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("Bye byee server")
	})

	http.ListenAndServe(":8000", nil)

}
