package handlers

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type hello struct {
	l *log.Logger
}

func NewHello(l *log.Logger) *hello {
	return &hello{l}
}

func (h *hello) ServeHTTP(rw http.ResponseWriter, r http.Request) {
	h.l.Println("Handle hello request")
	d, err := io.ReadAll(r.Body)

	if err != nil {
		h.l.Println("", err)
		http.Error(rw, "unable to read request body", http.StatusBadRequest)
		return
	}
	fmt.Println(rw, "hello is %s", d)

}
