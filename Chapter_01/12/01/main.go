package main

import (
	"01/packages/store"
	"fmt"
)

func main() {
	product := store.Product {
		Name: "Kayak",
		Category: "Water Sport",
	}

	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
}