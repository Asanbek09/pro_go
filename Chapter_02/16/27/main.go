package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := regexp.MustCompile("A ([A-z]*) for ([A-z]*) person")
	description := "Kayak. A boat for two persons"
	subs := pattern.FindStringSubmatch(description)

	for _, s := range subs {
		fmt.Println("Match:", s)
	}
}