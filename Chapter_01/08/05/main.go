package main

import "fmt"

func printSuppliers(product string, suppliers []string) {
	for _, supplier := range suppliers {
		fmt.Println("Product:", product, "Supplier:", supplier)
	}
}

func main() {
	printSuppliers("kayak", []string {"Acme kayak", "Bob's boats", "Crazy Canoe"})
	printSuppliers("Jacket", []string {"Saile safe Co"})
}