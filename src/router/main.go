package router

import (
	"go-backend/src/controller"
	"net/http"
)

func Router() {
	http.HandleFunc("/products", controller.ProductsController)
	http.HandleFunc("/products/", controller.ProductController)
}
