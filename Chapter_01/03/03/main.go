package main

import "fmt"

func main() {
	PrintHello()
	for i := 0; i < 5; i++ {
		i = i
		PrintNumber(i)
	}
}

func PrintHello() {
	fmt.Println("Hello, GO with vet example")
}

func PrintNumber(number int) {
	fmt.Println(number)
}
