package main

import "fmt"

func main() {
	price, tax, inStock := 250.00, 264.50, false
	fmt.Println("Total: ", price+tax)
	fmt.Println("In stock: ", inStock)

	// var price2, tax = 200.00, 24.05  ./main.go:6:9: other declaration of tax
	price2, tax := 200.00, 24.05
	fmt.Println("Total 2 : ", price2 + tax)
}
