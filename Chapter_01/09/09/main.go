package main

import "fmt"

type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func main() {
	watersportsProducts := map[string]float64 {
		"kayak": 275,
		"jacket": 48.95,
	}

	soccerProducts := map[string]float64 {
		"ball": 19.50,
		"stadium": 79500,
	}

	calc := func(price float64) float64 {
		if (price > 100) {
			return price + (price * 0.2)
		}
		return price;
	}

	for product, price := range watersportsProducts {
		printPrice(product, price, calc)
	}

	calc = func(price float64) float64 {
		if (price > 50) {
			return price + (price * 0.2)
		}
		return price
	}
	for product, price := range soccerProducts {
		printPrice(product, price, calc)
	}
}