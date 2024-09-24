package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "0b1100100"

	int1, int1err := strconv.ParseInt(val1, 0, 8)

	if int1err == nil {
		smallInt := int8(int1)
		fmt.Println("Parsed value:", smallInt)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}