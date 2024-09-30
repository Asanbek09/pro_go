package main

import "fmt"

func main() {
	products := []string {"kayak", "jacket", "paddle", "hat"}

	for index, value := range products[2:] {
		fmt.Println("Index:", index, "value:", value)
	}
}