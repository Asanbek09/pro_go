package main

import (
	"fmt"
	"strconv"
)

func main() {
	val := 49.95

	fString := strconv.FormatFloat(val, 'f', 2, 64)
	eString := strconv.FormatFloat(val, 'e', -1, 64)

	fmt.Println("Format F: " + fString)
	fmt.Println("Format E: " + eString)
}
