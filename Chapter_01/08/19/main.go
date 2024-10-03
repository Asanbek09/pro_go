package main

import "fmt"

func calcTax(price float64) (float64, bool) {
	if (price > 100) {
		return price * 0.2, true
	}
	return 0, false
}

func calcTotalPrice(products map[string]float64, minSpend float64) (total, tax float64) {
	total = minSpend
	for _, price := range products {
		if taxAmount, due := calcTax(price); due {
			total += taxAmount;
			tax += taxAmount
		} else {
			total += price
		}
	}
	return
}

func main() {
	products := map[string]float64 {
		"kayak": 275, "jacket": 48.95,
	}

	total1, tax1 := calcTotalPrice(products, 10)
	fmt.Println("Total 1:", total1, "Tax:", tax1)
	total2, tax2 := calcTotalPrice(nil, 10)
	fmt.Println("Total 2:", total2, "Tax:", tax2)
}