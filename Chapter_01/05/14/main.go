package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "100"

	int1, int1err := strconv.ParseInt(val1, 0, 8) 

	if int1err == nil {
		fmt.Println("Parsed value:", int1)
	} else {
		fmt.Println("Cannot parse", val1)
	}
}