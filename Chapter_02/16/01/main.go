package main

import (
	"fmt"
	"strings"
)

func main() {
	product := "kayak"

	fmt.Println("Contains:", strings.Contains(product, "yak"))
	fmt.Println("ContainsAny:", strings.ContainsAny(product, "abc"))
	fmt.Println("ContainsRune:", strings.ContainsRune(product, 'k'))
	fmt.Println("EqualFold:", strings.EqualFold(product, "KAYAK"))
	fmt.Println("HasPrefix:", strings.HasPrefix(product, "ka"))
	fmt.Println("HasSuffix:", strings.HasSuffix(product, "yak"))
}