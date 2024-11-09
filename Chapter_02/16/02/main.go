package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "a boat for sailing"

	fmt.Println("Original:", description)
	fmt.Println("Title:", strings.Title(description))
}