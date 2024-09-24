package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "100"

	int1, int1err := strconv.Atoi(val1)

	if int1err == nil {
		var intResult int = int1
		fmt.Println("Parsed value:", intResult)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}