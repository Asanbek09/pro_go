package main

import (
	"fmt"
	"strconv"
)

func main() {
	var price string = "$48.95"

	//var currency byte = price[0] // result = 36
	var currency string = string(price[0])
	var amountString string = price[1:]
	amount, parseErr := strconv.ParseFloat(amountString, 64)

	fmt.Println("length:", len(price))
	fmt.Println("Currency: ", currency)
	if (parseErr == nil) {
		fmt.Println("Amount:", amount)
	} else {
		fmt.Println("Parse error:", parseErr)
	}
}