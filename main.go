package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var products = []product{
	product{1, "baju", 100000, 90},
	product{2, "kemeja", 10000, 20},
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/products", productController)
	fmt.Println("Server is listenging on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello From Go!")
}

// list / get data
func productController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, _ := json.Marshal(products)
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}
