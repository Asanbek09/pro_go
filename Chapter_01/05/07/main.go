package main

import "fmt"

func main() {
	first := 100

	second := &first
	third := &first

	alpha := 100
	beta := &alpha

	fmt.Println(second == third) // сравнение области памяти
	fmt.Println(second == beta)  // сравнение области памяти

	fmt.Println(*second == *third) // сравнение значения переменной
	fmt.Println(*second == *beta)  // сравнение значения переменной
}
