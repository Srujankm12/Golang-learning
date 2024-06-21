package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type Hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}
func (h *Hello) ServeHTTP(rw http.ResponseWriter, r http.Request) {
	log.Println("server has started running")

	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "hey", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "hello %s", d)
	log.Println(string(d))
}
