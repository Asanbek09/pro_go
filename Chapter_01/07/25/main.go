package main

import "fmt"

func main() {
	products := [4]string {"kayak", "jacket", "paddle", "hat"}

	deleted := append(products[:2], products[3:]...)
	fmt.Println("Deleted:", deleted)
}