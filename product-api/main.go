package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/amankapur007/product-api/handlers"
)

func main() {
	fmt.Println("product api")
	l := log.New(os.Stdout, "products-api", log.LstdFlags)
	productsHandler := handlers.NewProducts(l)

	sm := http.NewServeMux()
	sm.Handle("/", productsHandler)
	s := http.Server{
		Addr:         ":9090",
		Handler:      sm,
		ErrorLog:     l,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 1 * time.Second,
		ReadTimeout:  1 * time.Second,
	}
	// start the server
	s.ListenAndServe()
}
