package main

import "fmt"

func main() {
	names := [3]string {"kayak", "jacket", "paddle"}
	moreNames := [3]string {"kayak", "jacket", "paddle"}

	same := names == moreNames

	fmt.Println("comparison:", same)
}