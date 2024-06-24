package handlers

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

func (H *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	H.l.Println("server has started")

	d, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error("error ", err)

		return
	}
	fmt.Fprintf(rw, "hello %s", d)
	log.Printf("data is %s", d)

}
