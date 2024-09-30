package main

import "fmt"

func main() {
	products := map[string]float64 {
		"kayak": 279, "jacket": 48.54, "hat": 0,
	}

	fmt.Println("Hat:", products["hat"])
}