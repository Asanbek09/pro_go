package main

import (
	"fmt"
	"strings"
)

func main() {
	description := "A boat for two people"

	splits := strings.SplitN(description, " ", 3)
	for _, x := range splits {
		fmt.Println("SplitN: >> " + x + "<<")
	}
}