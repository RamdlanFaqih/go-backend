package main

import (
	"fmt"
	"go-backend/src/router"
	"net/http"
)

func main() {
	router.Router()
	http.HandleFunc("/", sayHello)
	fmt.Println("Server is listenging on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello From Go!")
}
