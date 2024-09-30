package main

import "fmt"

func main() {
	products := [4]string {"kayak", "jacket", "paddle", "hat"}

	allNames := products[1:]
	someNames := []string {"boots", "canoe"}
	copy(someNames[1:], allNames[2:3])

	fmt.Println("someNames:", someNames)
	fmt.Println("allNames:", allNames)
}