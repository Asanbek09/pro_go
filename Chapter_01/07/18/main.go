package main

import "fmt"

func main() {
	products := [4]string {"kayak", "jacket", "paddle", "hat"}

	someNames := products[1:3:3]
	allNames := products[:]

	someNames = append(someNames, "gloves")
	//someNames = append(someNames, "boots")

	fmt.Println("someNames:", someNames)
	fmt.Println("someNames len:", len(someNames), "cap:", cap(someNames))
	fmt.Println("allNames:", allNames)
	fmt.Println("allNames len:", len(allNames), "cap:", cap(allNames))
}