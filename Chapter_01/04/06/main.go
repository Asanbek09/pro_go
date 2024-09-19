package main

import "fmt"

func main() {
	const price, tax float32 = 275, 254.90
	const quantity, inStock = 2, true
	fmt.Println("Total: ", quantity*(price+tax))
	fmt.Println("In Stock: ", inStock)
}
