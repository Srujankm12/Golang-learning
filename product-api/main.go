package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	l := log.New(os.Stdout, "prouct-api", log.LstdFlags)
	ph := handlers.NewProduct(l)
	sm := http.NewServeMux()

	sm.Handle("/", ph)

}
