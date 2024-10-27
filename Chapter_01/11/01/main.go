package main

import "fmt"

type Product struct {
	name, category string
	price float64
}

func main() {
	products := []*Product {
		{"kayak", "water sport", 215},
		{"ball", "soccer", 78.50},
		{"gas", "car", 67.90},
	}

	for _, p := range products {
		fmt.Println("Name:", p.name, "Category:", p.category, "Price", p.price)
	}
}