package main

import "fmt"

type Product struct {
	name, category string
	price float64
}

func newProduct(name, category string, price float64) *Product {
	return &Product{name, category, price}
}

func(product *Product) printDetails() {
	fmt.Println("Name:", product.name, "Category:", product.category, "Price:", product.price)
}

func main() {
	products := []*Product {
		{"Kayak", "Water sport", 285},
		{"Ball", "Soccer", 45.9},
		{"Motor", "Car", 1500},
	}

	for _, p := range products {
		p.printDetails()
	}
}