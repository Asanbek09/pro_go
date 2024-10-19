package main

import "fmt"

func calcWithTax(price float64) float64 {
	return price + (price * 0.2)
}

func calcWithoutTax(price float64) float64 {
	return price
}

func printPrice(product string, price float64, calculator func(float64) float64) {
	fmt.Println("Produc: ", product, "Price: ", calculator(price))
}

func selectCalculator(price float64) func(float64) float64 {
	if(price > 100) {
		return calcWithTax
	}
	return calcWithoutTax
}

func main() {
	products := map[string]float64 {
		"kayak": 275,
		"jacket": 48.95,
	}

	for product, price := range products {
		printPrice(product, price, selectCalculator(price))
	}
}