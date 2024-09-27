package main

import "fmt"

func main() {
	product := "Kayak"

	for index := range product {
		fmt.Println("Index:", index)
	}	
}