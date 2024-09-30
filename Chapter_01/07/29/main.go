package main

import "fmt"

func main() {
	p1 := []string {"kayak", "jacket", "paddle", "hat"}
	arrayPtr := (*[3]string)(p1)
	array := *arrayPtr

	fmt.Println(array)
}