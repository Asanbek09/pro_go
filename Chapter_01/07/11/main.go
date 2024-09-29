package main

import "fmt"

func main() {
	names := []string {"kayak", "jacket", "paddle"}

	names = append(names, "hat", "gloves")
	fmt.Println(names)
}