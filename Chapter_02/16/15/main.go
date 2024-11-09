package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat is for two persons"

	trimmer := func(r rune) bool {
		return r == 'A' || r == 's'
	}

	trimmed := strings.TrimFunc(description, trimmer)
	fmt.Println("Trimmed:", trimmed)
}