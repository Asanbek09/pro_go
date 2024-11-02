package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

type Person struct {
	name, city string
}

func processItem(item interface{}) {
	switch value := item.(type) {
	case Product:
		fmt.Println("Product:", value.name, "Price:", value.price)
	case *Product:
		fmt.Println("Product Pointer:", value.name, "Price:", value.price)
	case Service:
		fmt.Println("Service:", value.description, "Price:", value.monthlyFee * float64(value.durationMonths))
	case Person:
		fmt.Println("Name:", value.name, "City:", value.city)
	case *Person:
		fmt.Println("Name Pointer:", value.name, "City:", value.city)
	case string, bool, int:
		fmt.Println("Built-in type:", value)
	default:
		fmt.Println("Default:", value)
	}
}

func main() {
	var expense Expense = &Product {"Kayak", "Water Sport", 245}

	data := []interface{} {
		expense,
		Product {"jacket", "clothes", 54.97},
		Service {"Boat Cover", 12, 83.45, []string{}},
		Person {"Alice", "Boston"},
		&Person {"Bob", "New York"},
		"book",
		78,
		false,
	}

	for _, item := range data{
		processItem(item)
	}
}