package main

import "fmt"

func main() {
	fmt.Println("Hello, operations")

	price, tax := 265.00, 27.30
	sum := price + tax
	difference := price - tax
	product := price * tax
	quotient := price / tax

	fmt.Println(sum)
	fmt.Println(difference)
	fmt.Println(product)
	fmt.Println(quotient)
}
