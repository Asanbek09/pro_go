package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := regexp.MustCompile("A [A-z]* for [A-z]* person")
	description := "Kayak. A boat for two persons"

	str := pattern.FindString(description)
	fmt.Println("Match:", str)
}