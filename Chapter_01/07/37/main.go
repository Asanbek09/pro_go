package main

import (
	"fmt"
	"sort"
)

func main() {
	products := map[string]float64{
		"kayak": 64, "jacket": 56.32, "hat": 0,
	}

	keys := make([]string, 0, len(products))
	for key, _ := range products {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println("key: ", key, "value: ", products[key])
	}
}
