package main

import "fmt"

func printPrice(product string, price, _ float64) {
	taxAmount := price * 0.25
	fmt.Println(product, "price:", price, "tax:", taxAmount)
}

func main() {
	printPrice("kayak", 274, 0.2)
	printPrice("jacket", 283, 0.3)
	printPrice("ball", 123.65, 0.2)
}