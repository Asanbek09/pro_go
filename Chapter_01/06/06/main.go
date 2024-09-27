package main

import (
	"fmt"
	"strconv"
)

func main() {
	priceString := "275"

	if kayak, err := strconv.Atoi(priceString); err == nil {
		fmt.Println("Price:", kayak)
	} else {
		fmt.Println("Error:", err)
	}
}