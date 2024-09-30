package main

import (
	"fmt"
	"reflect"
)

func main() {
	p1 := []string {"kayak", "jacket", "paddle", "hat"}
	p2 := p1

	fmt.Println("Equal:", reflect.DeepEqual(p1, p2))
	// fmt.Println("Equal:", p1 == p2) // result invalid operation: p1 == p2 (slice can only be compared to nil)
}