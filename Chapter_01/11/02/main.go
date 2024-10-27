package main

import "fmt"

type Product struct {
	name, category string
	price float64
}

func printDetails(product *Product) {
	fmt.Println("Name:", product.name, "Category:", product.category, "Price:", product.price)
}

func main() {
	products := []*Product {
		{"kayak", "water sport", 255},
		{"Ball", "soccer", 45.09},
		{"Motor", "car", 15800},
	}

	for _, p := range products {
		printDetails(p)
	}
}