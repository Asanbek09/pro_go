package main

import "fmt"

func main() {
	var price = 275.00
	var tax float32 = 27.80
	fmt.Println(price + float64(tax)) // price + tax (mismatched types float64 and float32)
}
