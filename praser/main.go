package main

import (
	//"log"
	"log"
	"net/http"
	"os"
	"praser/handlers"
)

func main() {

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

	// })
	// http.HandleFunc("/close", func(rw http.ResponseWriter, r *http.Request) {
	// 	log.Println("server has stoped")

	// })
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":8001", sm)
}
