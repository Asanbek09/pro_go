package main

import "fmt"

func printPrice() {
	kayak := 275.00
	tax := kayak * 0.2
	fmt.Println("price:", kayak, "tax:", tax)
}

func main() {
	fmt.Println("About to call function")
	printPrice()
	fmt.Println("Function complete")
}