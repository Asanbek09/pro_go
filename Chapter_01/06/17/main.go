package main

import "fmt"

func main() {
	product := "Kayak"

	for index, character := range product {
		switch (character) {
			case 'K', 'k':
				fmt.Println("K or k at position", index)
			case 'y':
				fmt.Println("y at position", index)
		}
	}
}