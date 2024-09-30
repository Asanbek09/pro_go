package main

import (
	"fmt"
	"sort"
)

func main() {
	products := []string {"kayak", "jacket", "paddle", "hat"}

	sort.Strings(products)

	for index, value := range products {
		fmt.Println("Index:", index, "Value:", value)
	}
}