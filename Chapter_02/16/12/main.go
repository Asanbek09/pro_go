package main

import (
	"fmt"
	"strings"
)

func main() {
	username := " Bob"
	trimmed := strings.TrimSpace(username)
	fmt.Println("Trimmed:", trimmed)
}