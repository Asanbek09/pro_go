package main

import "fmt"

func main() {
	product := "Kayak"

	for index, character := range product {
		switch(character) {
			case 'K':
				fmt.Println("K at position", index)
			case 'y':
				fmt.Println("Y at position", index)
		}
	}
}