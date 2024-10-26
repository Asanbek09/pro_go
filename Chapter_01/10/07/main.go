package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price float64
	}

	p1 := Product {
		name: "kayak", category: "water sport", price: 250.00,
	}
	p2 := Product {
		name: "kayak", category: "water sport", price: 250.00,
	}
	p3 := Product {
		name: "kayak", category: "soccer", price: 240.00,
	}

	fmt.Println("p1 == p2:", p1 == p2)
	fmt.Println("p1 == p3:", p1 == p3)
}