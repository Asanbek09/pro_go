package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price float64
	}

	kayak := Product {
		name: "kayak",
		category: "water sports",
	}

	fmt.Println(kayak.name, kayak.category, kayak.price)
	kayak.price = 300
	fmt.Println("Changed price:", kayak.price)

	var jacket Product
	fmt.Println("Name is zero value:", jacket.name == "")
	fmt.Println("Category is zero value:", jacket.category == "")
	fmt.Println("Price is zero value:", jacket.price == 0)
}