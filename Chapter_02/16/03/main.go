package main

import (
	"fmt"
	"strings"
)

func main() {
	specialChar := "\u01c9"
	fmt.Println("Original:", specialChar, []byte(specialChar))

	upperChar := strings.ToUpper(specialChar)
	fmt.Println("Upper:", upperChar, []byte(upperChar))

	titleChar := strings.ToTitle(specialChar)
	fmt.Println("Title:", titleChar, []byte(titleChar))
}