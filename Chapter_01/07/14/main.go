package main

import "fmt"

func main() {
	names := make([]string, 3, 6)

	names[0] = "kayak"
	names[1] = "jacket"
	names[2] = "paddle"

	appendedNames := append(names, "hat", "gloves")

	names[0] = "canoe"

	fmt.Println("names:", names)
	fmt.Println("appendedNames:", appendedNames)
}