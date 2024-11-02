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

func getProducts() []Product {
	return []Product {
		{"kayak", "water sport", 290},
		{"jacket", "clothes", 190.55},
		{"ball", "soccer", 54.95},
	}
}

func main() {
	products := ProductList(getProducts())
	for category, total := range products.calcCategoryTotals() {
		fmt.Println("Category:", category, "Total:", total)
	}
}