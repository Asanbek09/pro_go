package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat is for two persons"
	trimmed := strings.Trim(description, "Asno ")
	fmt.Println("Trimmed:", trimmed)
}