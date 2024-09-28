package main

import "fmt"

func main() {
	names := [3]string{"kayak", "jacket", "paddle"}

	var otherArray [4]string = names

	fmt.Println(names)
}

// ./main.go:8:6: otherArray declared and not used
// ./main.go:8:29: cannot use names (variable of type [3]string) as [4]string value in variable declaration
