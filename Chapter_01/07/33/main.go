package main

import "fmt"

func main() {
	products := map[string]float64 {
		"kayak": 274, "jacket": 45.64, "hat": 0,
	}

	value, ok := products["hat"]

	if(ok) {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}
}