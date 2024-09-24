package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "48.95"

	float1, float1err := strconv.ParseFloat(val1, 64)

	if float1err == nil {
		fmt.Println("Parsed value:", float1)
	} else {
		fmt.Println("Cannot parse", val1, float1err)
	}
}