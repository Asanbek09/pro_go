package main

import (
	"fmt"
	"02/packages/store"
)

func main() {
	product := store.NewProduct("Kayak", "Water sport", 246)

	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	fmt.Println("Price:", product.Price())
}