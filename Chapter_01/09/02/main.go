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
		"kayak": 255,
		"jacket": 46.95,
	}

	for product, price := range products {
		var calcFunc func(float64) float64
		fmt.Println("Function assigned:", calcFunc == nil)
		if (price > 100) {
			calcFunc = calcWithTax
		} else {
			calcFunc = calcWithoutTax
		}
		fmt.Println("Function assigned:", calcFunc == nil)
		totalPrice := calcFunc(price)
		fmt.Println("Product:", product, "Total Price:", totalPrice)
	}
}