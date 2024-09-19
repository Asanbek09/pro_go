package main

import "fmt"

func main() {
	price, tax, inStock := 266.00, 27.40, true
	fmt.Println("Total: ", price + tax)
	fmt.Println("In Stock: ", inStock)
}