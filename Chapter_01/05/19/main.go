package main

import (
	"fmt"
	"strconv"
)

func main() {
	val1 := "100"

	int1, int1err := strconv.ParseInt(val1, 10, 0)
	
	if int1err == nil {
		var intResult int = int(int1)
		fmt.Println("Parsed value:", intResult)
	} else {
		fmt.Println("Cannot parse", val1, int1err)
	}
}