package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price float64
	}

	var kayak = Product {"kayak", "water sport", 275.05}

	fmt.Println("Name:", kayak.name)
	fmt.Println("Category:", kayak.category)
	fmt.Println("Price:", kayak.price)
}