package main

import "fmt"

func printPrice(product string, price, tax float64) {
	taxAmount := price * tax
	fmt.Println(product, "price:", price, "tax:", taxAmount)
}

func main() {
	printPrice("kayak", 275, 0.2)
	printPrice("jacket", 48.95, 0.1)
	printPrice("ball", 19.43, 0.5)
}