package main

import "fmt"

func printPrice(product string, price float64, taxRate float64) {
	taxAmount := price * taxRate
	fmt.Println(product, "price:", price, "Tax:", taxAmount)
}

func main() {
	printPrice("kayak", 275, 0.2)
	printPrice("jacket", 48.95, 0.2)
	printPrice("ball", 19.50, 0.15)
}