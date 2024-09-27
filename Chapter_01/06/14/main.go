package main

import "fmt"

func main() {
	product := "Product"

	for _, character := range product {
		fmt.Println("Character:", string(character))
	}
}