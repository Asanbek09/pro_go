package main

import "fmt"

func printSuppliers(product string, suppliers ...string) {
	if (len(suppliers) == 0) {
		fmt.Println("Product:", product, "Supplier: (none)")
	} else {
		for _, supllier := range suppliers {
			fmt.Println("Product:", product, "Supplier:", supllier)
		}
	}
}

func main() {
	printSuppliers("kayak", "acme kayak", "bob boats", "crazy canoe")
	printSuppliers("jacket", "Sail safe Co")
	printSuppliers("Soccer Ball")
}