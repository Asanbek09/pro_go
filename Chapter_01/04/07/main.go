package main

import "fmt"

func main() {
	var price float32 = 275.00
	var tax float32 = 27.70
	fmt.Println(price + tax)
	price = 300
	fmt.Println(price + tax)
}
