// models/product.go
package models

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Stock int     `json:"stock"`
}

var Products = []Product{
	{ID: "1", Name: "Laptop", Price: 999.99, Stock: 100},
	{ID: "2", Name: "Phone", Price: 499.99, Stock: 200},
}
