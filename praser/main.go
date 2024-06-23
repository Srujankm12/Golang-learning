package main

import (
	//"log"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"praser/handlers"
	"time"
)

func main() {

	// http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {

	// })
	// http.HandleFunc("/close", func(rw http.ResponseWriter, r *http.Request) {
	// 	log.Println("server has stoped")

	// })
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)
	gb := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gb)

	s := http.Server{
		Addr:         ":8001",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	//http.ListenAndServe(":8001", sm)
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}

	}()
	sigcha := make(chan os.Signal)
	signal.Notify(sigcha, os.Interrupt)
	signal.Notify(sigcha, os.Kill)

	sig := <-sigcha
	l.Println("Recived terminate ,gracefull shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
