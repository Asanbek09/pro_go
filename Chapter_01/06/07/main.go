package main

import "fmt"

func main() {
	counter := 0
	for {
		fmt.Println("Counter:", counter)
		counter++
		if (counter > 3) {
			break
		}
	}
}