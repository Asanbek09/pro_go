package main

import "fmt"

func main() {
	products := [4]string {"kayak", "jacket", "paddle", "hat"}

	allNames := products[1:]
	someNames := make([]string, 2)
	copy(someNames, allNames)

	fmt.Println("someNames:", someNames)
	fmt.Println("allNames:", allNames)
}