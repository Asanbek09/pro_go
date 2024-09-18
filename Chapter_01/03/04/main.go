package main

import "fmt"

func main() {
	PrintHello()

	for i := 0; i < 5; i++ { //loop with a counter
		PrintHello()   //print out message
		PrintNumber(i) //print out the counter
	}
}

func PrintHello() {
	fmt.Println("Hello, Go working with go fmt")
}

func PrintNumber(number int) {
	fmt.Println(number)
}
