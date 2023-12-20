package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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
	http.HandleFunc("/products", productsController)
	http.HandleFunc("/products/", productController)
	fmt.Println("Server is listenging on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello From Go!")
}

// list / get data
func productsController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		result, _ := json.Marshal(products)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if r.Method == "POST" {
		var product product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid request body")
			return
		}
		products = append(products, product)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "product created successfully")
	}
	http.Error(w, "", http.StatusBadRequest)
}

// get detail by id
func productController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParams := r.URL.Path[len("/products/"):]
	id, _ := strconv.Atoi(idParams)
	// fmt.Fprintln(w, reflect.TypeOf(id))

	var foundIndex = -1
	for i, p := range products {
		if p.Id == id {
			foundIndex = i
			break
		}
	}

	if foundIndex == -1 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	result, _ := json.Marshal(products[foundIndex])
	w.WriteHeader(http.StatusOK)
	w.Write(result)
}

// create product
