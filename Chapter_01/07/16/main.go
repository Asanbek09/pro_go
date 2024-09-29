package main

import "fmt"

func main() {
	products := [4]string {"kayak", "jacket", "paddle", "hat"}

	someNames := products[1:3]
	allNames := products[:]

	fmt.Println("someNames:", someNames)
	fmt.Println("allNames:", allNames)
}