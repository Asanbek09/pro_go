package main

import "fmt"

func main() {
	for counter := 0; counter <= 3; counter++ {
		if (counter == 1) {
			continue
		}
		fmt.Println("Counter:", counter)
	}
}