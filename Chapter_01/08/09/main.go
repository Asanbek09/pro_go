package main

import "fmt"

func printSuppliers(product string, suppliers ...string) {
	if (len(suppliers) == 0) {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supplier := range suppliers {
			fmt.Println("Product:", product, "Supplier:", supplier)
		}
	}
}

func main() {
	names := []string {"Acme Kayak", "Bob's Boats", "Crazy Canoe"}

	printSuppliers("kayak", names...)
	printSuppliers("Jacket", "Sail safe Co")
	printSuppliers("Soccer Ball")
}