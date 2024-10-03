package main

import "fmt"

func calcTax(price float64) (float64, bool) {
	if (price > 100) {
		return price * 0.2, true
	}
	return 0, false
}

func main() {
	products := map[string]float64 {
		"kayak": 287.75,
		"jacket": 48.55,
	}

	for product, price := range products {
		taxAmount, taxDue := calcTax(price)
		if (taxDue) {
			fmt.Println("Product:", product, "Price:", taxAmount)
		} else {
			fmt.Println("Product:", product, "No tax due")
		}
	}
}