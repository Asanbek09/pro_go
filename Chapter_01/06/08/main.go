package main

import "fmt"

func main() {
	counter := 0
	for (counter <= 3) {
		fmt.Println("Counter:", counter)
		counter++
	}
}