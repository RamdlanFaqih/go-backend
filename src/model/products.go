package model

type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

var Products = []Product{
	Product{1, "baju", 100000, 90},
	Product{2, "kemeja", 10000, 20},
}
