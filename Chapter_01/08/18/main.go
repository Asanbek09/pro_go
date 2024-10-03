package main

import "fmt"

func calcTotalPrice(products map[string]float64) (count int, total float64) {
	count = len(products)
	for _, price := range products {
		total += price
	}
	return
}

func main() {
	products := map[string]float64 {
		"kayak": 275, "jacket": 43.95,
	}

	_, total := calcTotalPrice(products)
	fmt.Println("Total:", total)
}