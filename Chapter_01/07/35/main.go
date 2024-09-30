package main

import "fmt"

func main() {
	products := map[string]float64 {
		"kayak": 42, "jacket": 94.5, "hat": 0,
	}

	delete(products, "hat")

	if value, ok := products["hat"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}
}