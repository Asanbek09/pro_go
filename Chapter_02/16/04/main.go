package main

import (
	"fmt"
	"unicode"
)

func main() {
	product := "Kayak"

	for _, p := range product {
		fmt.Println(string(p), "Upper case:", unicode.IsUpper(p))
	}
}