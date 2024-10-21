package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price float64
	}

	type StockLevel struct {
		Product
		Alternate Product
		count int
	}

	stockItem := StockLevel {
		Product: Product {"kayak", "water sport", 275.50},
		Alternate: Product{"jacket", "water sports", 35.98},
		count: 50,
	}

	fmt.Println("Name:", stockItem.Product.name)
	fmt.Println("Count", stockItem.count)
	fmt.Println("Alt name:", stockItem.Alternate.name)
}