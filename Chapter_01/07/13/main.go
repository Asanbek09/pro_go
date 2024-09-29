package main

import "fmt"

func main() {
	names := make([]string, 3, 6)

	names[0] = "kayak"
	names[1] = "jacket"
	names[2] = "paddle"

	fmt.Println("len:", len(names))
	fmt.Println("cap:", cap(names))
}