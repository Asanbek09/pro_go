package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	kayak := Product {"Kayak", "water sport", 255}
	insurance := Service {"Boat Cover", 18, 99.55}

	fmt.Println("Product:", kayak.name, "Price:", kayak.price)
	fmt.Println("Service:", insurance.description, "Price:", insurance.monthlyFee * float64(insurance.durationMonths))
}