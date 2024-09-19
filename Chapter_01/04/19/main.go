package main

import "fmt"

func main() {
	first := 100
	second := &first

	first++
	*second++

	fmt.Println("First: ", first)
	fmt.Println("Second: ", *second)
}
