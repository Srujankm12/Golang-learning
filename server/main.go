package main

import (
	"fmt"
	"io"

	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		log.Println("server has started running")

		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, "hey", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "hello %s", d)
		log.Println(string(d))

	})

	http.HandleFunc("/bye", func(http.ResponseWriter, *http.Request) {
		log.Println("Bye byee server")
	})

	http.ListenAndServe(":8000", nil)

}
