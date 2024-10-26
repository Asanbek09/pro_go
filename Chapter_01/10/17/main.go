package main

import "fmt"

type Product struct {
	name, category string
	price float64
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func main() {
	products := [2]*Product {
		newProduct("kayak", "water sport", 265),
		newProduct("hat", "skiing", 150),
	}

	for _, p := range products {
		fmt.Println("Name", p.name, "Category:", p.category, "Price:", p.price)
	}
}