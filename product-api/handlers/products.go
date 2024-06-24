package handlers

import (
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle Product Request")
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Product API"))
}
