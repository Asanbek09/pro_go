package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat is for two persons"

	prefixTrimmed := strings.TrimPrefix(description, "A boat")
	wrongPrefix := strings.TrimPrefix(description, "A cap")

	fmt.Println("Prefix Trimmed:", prefixTrimmed)
	fmt.Println("Wrong Prefix:", wrongPrefix)
}