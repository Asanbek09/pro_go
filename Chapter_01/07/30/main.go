package main

import "fmt"

func main() {
	products := make(map[string]float64, 10)
	products["kayak"] = 279
	products["jacket"] = 48.95

	fmt.Println("Map size:", len(products))
	fmt.Println("Price:", products["kayak"])
	fmt.Println("Price:", products["hat"])
}