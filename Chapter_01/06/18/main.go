package main

import "fmt"

func main() {
	product := "kayakK"

	for index, character := range product {
		switch (character) {
		case 'K', 'k':
			if(character == 'k') {
				fmt.Println("Lowercase k at position", index)
				break
			}
			fmt.Println("Uppercase K at position", index)
		case 'y':
			fmt.Println("y at position", index)
		}
	}
}