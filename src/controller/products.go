package controller

import (
	"encoding/json"
	"fmt"
	"go-backend/src/model"
	"net/http"
	"strconv"
)

// list data & post
func ProductsController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		result, _ := json.Marshal(model.Products)
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if r.Method == "POST" {
		var product model.Product
		err := json.NewDecoder(r.Body).Decode(&product)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid request body")
			return
		}
		model.Products = append(model.Products, product)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintln(w, "product created successfully")
	}
	http.Error(w, "", http.StatusBadRequest)
}

// get detail by id, update, and delete
func ProductController(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idParams := r.URL.Path[len("/products/"):]
	id, _ := strconv.Atoi(idParams)
	// fmt.Fprintln(w, reflect.TypeOf(id))
	var foundIndex = -1
	for i, p := range model.Products {
		if p.Id == id {
			foundIndex = i
			break
		}
	}
	if foundIndex == -1 {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	if r.Method == "GET" {
		result, _ := json.Marshal(model.Products[foundIndex])
		w.WriteHeader(http.StatusOK)
		w.Write(result)
	} else if r.Method == "PUT" {
		var updateProduct model.Product
		err := json.NewDecoder(r.Body).Decode(&updateProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "invalid request body")
			return
		}

		model.Products[foundIndex] = updateProduct
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, "Product updated successfully")
		return
	} else if r.Method == "DELETE" {
		model.Products = append(model.Products[:foundIndex], model.Products[foundIndex+1:]...)
		w.WriteHeader(http.StatusAccepted)
		fmt.Fprintln(w, "Product Deleted")
	}
	http.Error(w, "", http.StatusBadRequest)
}
