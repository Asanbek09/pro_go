package main

import "fmt"

func main() {
	kayakPrice := 275.00

	if (kayakPrice > 500) {
		fmt.Println("Price is greater than 500")
	} else if (kayakPrice < 100) {
		fmt.Println("Price is less than 100")
	} else {
		fmt.Println("Price not matched by earlier expressions")
	}
}