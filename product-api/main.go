package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("product api")
	// start the server
	http.ListenAndServe(":9090", nil)
}
