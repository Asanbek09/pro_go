package main

import (
	"fmt"
	currencyFmt "03/packages/fmt"     // currencyFmt - это псевдоним
	"03/packages/store"
	"03/packages/store/cart"
	_ "03/packages/data"
	//. "03/packages/fmt" // Существует специальный псевдоним, известный как точечный импорт, который позволяет использовать функции пакета без использования префикса
)

func main() {
	product := store.NewProduct("Kayak", "Water Sport", 290)

	cart := cart.Cart {
		CustomerName: "Alice",
		Products: []store.Product{ *product},
	}

	fmt.Println("Name:", cart.CustomerName)
	fmt.Println("Total:", currencyFmt.ToCurrency(cart.GetTotal()))

	/*
	fmt.Println("Name:", product.Name)
	fmt.Println("Category:", product.Category)
	//fmt.Println("Price:", product.Price())
	fmt.Println("Price:", currencyFmt.ToCurrency(product.Price()))
	*/
}