package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) persons")

	description := "Kayak. A boat for two persons"

	template := "(type: ${type}, capacity: ${capacity})"
	replaced := pattern.ReplaceAllString(description, template)
	fmt.Println("Replaced:", replaced)
}