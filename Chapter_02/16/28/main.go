package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := regexp.MustCompile("A (?P<type>[A-z]*) for (?P<capacity>[A-z]*) person")

	description := "Kayak. A boat for two persons"

	subs := pattern.FindStringSubmatch(description)
	for _, name := range []string {"type", "capacity"} {
		fmt.Println(name, "=", subs[pattern.SubexpIndex(name)])
	}
}