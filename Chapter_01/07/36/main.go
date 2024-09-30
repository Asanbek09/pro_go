package main

import "fmt"

func main() {
	products := map[string]float64 {
		"kayak": 54.67, "jacket": 54, "hat": 0,
	}

	for key, value := range products {
		fmt.Println("Key", key, "Value:", value)
	}
}