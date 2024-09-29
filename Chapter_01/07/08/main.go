package main

import "fmt"

func main() {
	names := [3]string {"kayak", "jacket", "paddle"}

	//for index, value := range names {
	//	fmt.Println("Index:", index, "value:", value)
	//}

	for _, value := range names {
		fmt.Println("Value:", value)
	}
}