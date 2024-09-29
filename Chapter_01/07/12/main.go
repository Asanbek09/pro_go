package main

import "fmt"

func main() {
	names := []string {"kayak", "jacket", "paddle"}

	appendedNames := append(names, "hat", "gloves")

	names[0] = "Canoe"

	fmt.Println("names:", names)
	fmt.Println("appendedNames:", appendedNames)
}