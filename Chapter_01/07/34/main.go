package main

import "fmt"

func main() {
	products := map[string]float64 {
		"kayak": 432, "jacket": 75.89, "hat":0,
	}

	if value, ok := products["hat"]; ok {
		fmt.Println("Stored value:", value)
	} else {
		fmt.Println("No stored value")
	}
}