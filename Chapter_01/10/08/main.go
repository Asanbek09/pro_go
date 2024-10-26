package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price float64
	}

	type Item struct {
		name, category string
		price float64
	}

	prod := Product {
		name: "kayak", category: "water sport", price: 275.00,
	}
	item := Item {
		name: "kayak", category: "water sport", price: 275.00,
	}

	fmt.Println("prod == item:", prod == Product(item))
}