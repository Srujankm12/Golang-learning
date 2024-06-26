package handlers

import (
	"log"
	"net/http"
	"newprac/data"
	"regexp"
	"strconv"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}
	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)

		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return

		}
		if len(g) != 1 {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return

		}
		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {

			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}
		p.updateProducts(id,rw http.ResponseWriter, r *http.Request)
		return
		



		// p.l.Println("got id", id)

	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal JSON", http.StatusInternalServerError)
	}
}
func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Products")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
	}
	p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)

}

func (p product)updateProducts(id int,rw http.ResponseWriter ,r*http.Request ){
	p.l.Println("handle Put product")

	prod:=&data.Product{}
	err :=prod.FromJSON(r.Body)
	if err!=nil{
		http.Error(rw,"unable to marshal json",http.StatusBadGateway)
	}
	data.updateProducts(id,prod)
}