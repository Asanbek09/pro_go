package main

import "fmt"

func main() {
	price, tax, inStock, _ := 274.00, 234.5, true, true
	var _ = "Alice"
	fmt.Println("Total: ", price + tax)
	fmt.Println("In stock: ", inStock)
}