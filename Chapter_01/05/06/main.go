package main

import "fmt"

func main() {
	fmt.Println("...")

	first := 100
	const second = 200.00

	equal := first == second
	notEqual := first != second
	lesssThan := first < second
	lesssThanOrEqual := first <= second
	greaterThan := first > second
	greaterThanOrEqual := first >= second

	fmt.Println(equal)
	fmt.Println(notEqual)
	fmt.Println(lesssThan)
	fmt.Println(lesssThanOrEqual)
	fmt.Println(greaterThan)
	fmt.Println(greaterThanOrEqual)
}
