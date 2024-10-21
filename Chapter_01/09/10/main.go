package main

import "fmt"

type calcFunc func(float64) float64

func printPrice(product string, price float64, calculator calcFunc) {
	fmt.Println("Product:", product, "Price:", calculator(price))
}

func priceCalcFactory(threshold, rate float64) calcFunc {
	return func(price float64) float64 {
		if (price > threshold) {
			return price + (price * rate)
		}
		return price
	}
}

func main() {
	watersportsProducts := map[string]float64 {
		"kayak": 275,
		"jacket": 48.95,
	}

	soccerProducts := map[string]float64 {
		"ball": 19.5,
		"stadium": 79500,
	}

	waterCalc := priceCalcFactory(100, 0.2);
	soccerCalc := priceCalcFactory(50, 0.1)

	for product, price := range watersportsProducts {
		printPrice(product, price, waterCalc)
	}

	for product, price := range soccerProducts {
		printPrice(product, price, soccerCalc)
	}
}