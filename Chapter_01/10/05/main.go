package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price float64
	}

	type StockLevel struct {
		Product
		count int
	}

	stockItem := StockLevel {
		Product: Product {"kayak", "water sport", 275.5},
		count: 40,
	}

	fmt.Println("Name:", stockItem.Product.name)
	fmt.Println("Count:", stockItem.count)
}