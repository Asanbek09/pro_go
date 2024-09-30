package main

import "fmt"

func main() {
	products := map[string]float64 {
		"kayak": 279, "jacket": 48.95,
	}

	fmt.Println("Map size:", len(products))
	fmt.Println("Price:", products["kayak"])
	fmt.Println("Price:", products["hat"])
}