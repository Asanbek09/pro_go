package main

import "fmt"

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
	return price
}

func main() {
	products := map[string]float64 {
		"kayak": 254,
		"jacket": 48.93,
	}

	for product, price := range products {
		var calcFunc func(float64) float64
		if (price > 100) {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}

		totalPrice := calcFunc(price)
		fmt.Println("Product:", product, "Price:", totalPrice)
	}
}