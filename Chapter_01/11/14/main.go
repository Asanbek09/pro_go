package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

type Person struct {
	name, city string
}

func main() {
	var expense Expense = &Product {"Kayak", "Water Sport", 175}

	data := []interface{} {
		expense,
		Product {"jacket", "clothes", 280.90},
		Service {"Boat Cover", 11, 85.99, []string{} },
		Person {"Alice", "Brandon"},
		&Person {"Bob", "New York"},
		"This is string",
		100,
		true,
	}

	for _, item := range data {
		switch value := item.(type) {
		case Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		case *Product:
			fmt.Println("Product Pointer:", value.name, "Price:", value.price)
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee * float64(value.durationMonths))
		case Person:
			fmt.Println("Person:", value.name, "City:", value.city)
		case *Person:
			fmt.Println("Person Pointer:", value.name, "City:", value.city)
		case string, bool, int:
			fmt.Println("Built-in type:", value)
		default:
			fmt.Println("Default:", value)
		}
	}
}