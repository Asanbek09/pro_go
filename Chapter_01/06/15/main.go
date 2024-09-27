package main

import "fmt"

func main() {
	products := []string {"Kayak", "Lifejacket", "soccerBall"}

	for index, element := range products {
		fmt.Println("Index", index, "Element:", element)
	} 
}