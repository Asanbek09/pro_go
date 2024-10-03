package main

import "fmt"

func calcTax(price float64) float64 {
	return price + (price * 0.2)
}

func main() {
	products := map[string]float64 {
		"kayak": 275, "jacket": 48.95,
	}

	for product, price := range products {
		priceWithTax := calcTax(price)
		fmt.Println("Product: ", product, "Price: ", priceWithTax)
	}
}