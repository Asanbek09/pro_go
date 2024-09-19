package main

import "fmt"

func main() {
	const prise float32 = 275.00
	const tax float32 = 54.20
	//const quantity int = 2
	const quantity = 2
	fmt.Println("Total: ", quantity*(prise+tax))
}
