package main

import (
	"fmt"
	"regexp"
)

func main() {
	pattern := regexp.MustCompile(" |boat|two")
	description := "A boat for two persons"

	split := pattern.Split(description, -1)

	for _, s := range split {
		if s != "" {
			fmt.Println("Substring:", s)
		}
	}
}