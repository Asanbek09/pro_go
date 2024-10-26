package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price float64
	}

	p1 := Product {
		name: "kayak",
		category: "water sport",
		price: 274.05,
	}

	p2 := p1
	p1.name = "original kayak"
	fmt.Println("P1:", p1.name)
	fmt.Println("P2:", p2.name)
}