package main

import "fmt"

func main() {
	first := 100
	var second *int

	// fmt.Println(*second) runtime error: invalid memory address or nil pointer dereference
	fmt.Println(second)
	second = &first
	fmt.Println(second)
}