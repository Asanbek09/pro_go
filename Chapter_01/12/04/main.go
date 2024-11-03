package main

import (
	"fmt"
	currencyFmt "04/packages/fmt"
	"04/packages/store"
	"04/packages/store/cart"
	_ "04/packages/data"
	//"github.com/fatih/color"
)

func main() {
	product := store.NewProduct("Kayak", "Water Sport", 290)

	cart := cart.Cart {
		CustomerName: "Alice",
		Products: []store.Product{ *product},
	}

	fmt.Println("Name:", cart.CustomerName)
	fmt.Println("Total:", currencyFmt.ToCurrency(cart.GetTotal()))
	//color.Green("Name:" + cart.CustomerName)
	//color.Cyan("Total:" + currencyFmt.ToCurrency(cart.GetTotal()))
}