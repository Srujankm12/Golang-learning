package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("server is running")

		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "error", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(rw, "Hello %s", d)

		log.Printf("data is %s", d)

	})
	http.HandleFunc("/end", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Close the server ")
	})
	log.Println("server is running")

	http.ListenAndServe(":8000", nil)

}
