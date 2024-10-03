package main

import "fmt"

func calcTax(price float64) float64 {
	if (price > 100) {
		return price * 0.2
	}
	return -1
}

func main() {
	products := map[string]float64 {
		"kayak": 275, "jacket": 48.95,
	}

	for product, price := range products {
		tax := calcTax(price)
		if (tax != -1) {
			fmt.Println("Product:", product, "Tax:", tax)
		} else {
			fmt.Println("Product:", product, "No tax due")
		}
	}
}