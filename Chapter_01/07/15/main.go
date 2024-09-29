package main

import "fmt"

func main() {
	names := make([]string, 3, 6)

	names[0] = "kayak"
	names[1] = "jacket"
	names[2] = "paddle"

	moreNames := []string {"Hat gloves"}

	appendedNames := append(names, moreNames...)
	fmt.Println("appendedNames", appendedNames)
}