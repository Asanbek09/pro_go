package main

import "fmt"

func writeName(val struct {
	name, category string
	price float64}) {
		fmt.Println("Name:", val.name)
	}

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
		name: "kayak", category: "water sport", price: 244.00,
	}
	item := Item {
		name: "stadium", category: "soccer", price: 89000,
	}

	writeName(prod)
	writeName(item)
}