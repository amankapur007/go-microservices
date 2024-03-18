package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/amankapur007/product-api/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l: l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("invalid id")
			http.Error(rw, "invalid url with id", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("invalid id")
			http.Error(rw, "invalid url with id", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(g[0][1])
		if err != nil {
			p.l.Println("unable to convert the id to int")
			http.Error(rw, "wrong id", http.StatusBadRequest)
			return
		}
		updateProduct(id, rw, r)
	}

	if r.Method == http.MethodDelete {
		reg := regexp.MustCompile(`/([0-9]+)`)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		if len(g) != 1 {
			p.l.Println("invalid id")
			http.Error(rw, "invalid url with id", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("invalid id")
			http.Error(rw, "invalid url with id", http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(g[0][1])
		if err != nil {
			p.l.Println("unable to convert the id to int")
			http.Error(rw, "wrong id", http.StatusBadRequest)
			return
		}
		deleteProduct(id, rw, r)
	}
}

func deleteProduct(id int, rw http.ResponseWriter, r *http.Request) {
	data.DeleteProductByID(id)
}

func updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to marshal the json", http.StatusBadRequest)
		return
	}
	err = data.UpdateProduct(id, product)
	if err != nil {
		http.Error(rw, "unable update the json", http.StatusBadRequest)
		return
	}

}

func getProducts(rw http.ResponseWriter, r *http.Request) {
	products := data.GetProducts()
	products.ToJSON(rw)
}

func addProduct(rw http.ResponseWriter, r *http.Request) {
	product := &data.Product{}
	err := product.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to marshal the json", http.StatusBadRequest)
		return
	}
	data.AddProduct(product)
}
