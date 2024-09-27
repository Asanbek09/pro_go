package main

import "fmt"

func main() {
	for counter := 0; true; counter++ {
		fmt.Println("Counter:", counter)
		if (counter > 3) {
			break
		}
	}
}