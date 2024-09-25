package main

import "fmt"

func main() {
	kayak := 275.00

	if (kayak > 500) {
		fmt.Println("Price is greater than 500")
	} else if (kayak < 300) {
		fmt.Println("Price is less than 300")
	}
}