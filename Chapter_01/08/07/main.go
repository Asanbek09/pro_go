package main

import "fmt"

func printSuppliers(product string, suppliers ...string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

func main() {
	printSuppliers("kayak", "acme kayaks", "bob's boats", "crazy canoe")
	printSuppliers("jacket", "sail safe Co")
	printSuppliers("soccer ball")
}