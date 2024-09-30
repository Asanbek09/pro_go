package main

import "fmt"

func main() {
	products := [4]string {"kayak", "jacket", "paddle", "hat"}

	allNames := products[1:]
	someNames := allNames[1:3]

	allNames = append(allNames, "gloves")
	allNames[1] = "canoe"

	fmt.Println("someNames:", someNames)
	fmt.Println("allNames:", allNames)
}