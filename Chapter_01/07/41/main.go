package main

import "fmt"

func main() {
	var price = "€48.95"

	for index, char := range []byte(price) {
		fmt.Println(index, char)
	}
}