package main

import "fmt"

type Product struct {
	name, category string
	price float64
}

type ProductList []Product

func (products *ProductList) calcCategoryTotals() map[string]float64 {
	totals := make(map[string]float64)
	for _, p := range *products {
		totals[p.category] = totals[p.category] + p.price
	}
	return totals
}

func main() {
	products := ProductList {
		{"kayak", "water sport", 265},
		{"jacket", "leather", 190.90},
		{"ball", "soccer", 75.45},
	}

	for category, total := range products.calcCategoryTotals() {
		fmt.Println("Category:", category, "Total:", total)
	}
}