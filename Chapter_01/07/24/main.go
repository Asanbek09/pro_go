package main

import "fmt"

func main() {
	products := []string {"kayak", "jacket", "paddle", "hat"}

	replacementProducts := []string {"canoe", "boots"}

	copy(products[0:1], replacementProducts)
	fmt.Println("products:", products)
}