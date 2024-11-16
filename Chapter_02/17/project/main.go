package main

import (
	"fmt"
)

func Printfln(template string, values ...interface{}) {
	fmt.Printf(template + "\n", values...)
}

/*
func getProductName(index int) (name string, err error) {
	if (len(Products) > index) {
		name = fmt.Sprintf("Name of product: %v", Products[index].Name)
	} else {
		err = fmt.Errorf("error for index %v", index)
	}
	return
}
*/

func main() {
	var name string
	var category string
	var price float64

	source := "Product jakcet watersports 264.59"
	template := "Product %s %s %f"
	n, err := fmt.Sscanf(source, template, &name, &category, &price)

	//fmt.Print("Enter text to scan: ")
	//n, err := fmt.Scanln(&name, &category, &price)

	//source := "Lifejacket watersports 264.50"
	//n, err := fmt.Sscan(source, &name, &category, &price)

	if(err == nil) {
		Printfln("Scanned %v values", n)
		Printfln("Name: %v, Category: %v, Price: %.2f", name, category, price)
	} else {
		Printfln("Error: %v", err.Error())
	}

/*
	name := "Kayak"
	Printfln("Bool: %t", len(name) > 1)
	Printfln("Bool: %t", len(name) > 100)
	Printfln("Pointer: %p", &name)
*/

/*
	name := "Kayak"
	Printfln("String: %s", name)
	Printfln("Charater: %c", []rune(name)[0])
	Printfln("Unicode: %U", []rune(name)[0])
*/

/*
	number := 279.00
	Printfln("Sign: >>%+.2f<<", number)
	Printfln("Zeros for Padding: >>%010.2f<<", number)
	Printfln("Right Reading: >>%-8.2f<<", number)
*/

/*
	Printfln("Decimalless with exponent: %b", number)
	Printfln("Decimal with exponent: %e", number)
	Printfln("Decimal without exponent: %f", number)
	Printfln("Hexadecimal: %x, %X", number, number)
	Printfln("Decimal without exponent: >>%8.2f<<", number)
	Printfln("Decimal without exponent: >>%.2f<<", number)
*/

/*
	Printfln("Binary: %b", number)
	Printfln("Decimal: %d", number)
	Printfln("Octal: %o, %O",number, number) 
	Printfln("Hexadecimal: %x, %X", number, number)
*/

/*
	Printfln("Value: %v", Kayak)
	Printfln("Value with fields: %+v", Kayak)
*/
	/*
	Printfln("Value: %v", Kayak)
	Printfln("Go syntax: %#v", Kayak)
	Printfln("Type: %T", Kayak)
	*/

	/*
	name, _ := getProductName(1)
	fmt.Println(name)

	_, err := getProductName(10)
	fmt.Println(err.Error())
	*/

	//fmt.Println("Product", Kayak.Name, "Price:", Kayak.Price)
	//fmt.Print("Product:", Kayak.Name, "Price:", Kayak.Price, "\n")
	//fmt.Printf("Product: %v, Price: $%4.2f", Kayak.Name, Kayak.Price)
}