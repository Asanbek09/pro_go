package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func main() {
	expenses := []Expense {
		Service {"Boat Cover", 12, 89.5, []string{}},
		Service {"Paddle Protect", 12, 8, []string{}},
		&Product {"Kayak", "Water Sport", 255},
	}

	for _, expense := range expenses {
		switch value := expense.(type) {
		case Service:
			fmt.Println("Service:", value.description, "Price:", value.monthlyFee * float64(value.durationMonths))
		case *Product:
			fmt.Println("Product:", value.name, "Price:", value.price)
		default:
			fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
		}
	}

	/*
	for _, expense := range expenses {
		if s, ok := expense.(Service); ok {
			fmt.Println("Service:", s.description, "Price:", s.monthlyFee * float64(s.durationMonths))
		} else {
			fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
		}
	}
	*/

	/*
	for _, expense := range expenses {
		s := expense.(Service)
		fmt.Println("Service:", s.description, "Price:", s.monthlyFee * float64(s.durationMonths)) 
	}
	*/
}

/*
func main() {
	var e1 Expense = &Product {name: "Kayak"}
	var e2 Expense = &Product {name: "Kayak"}
	var e3 Expense = Service {description: "Boat Cover"}
	var e4 Expense = Service {description: "Boat Cover"}

	fmt.Println("e1 == e2", e1 == e2)
	fmt.Println("e3 == e4", e3 == e4)
}
*/

/*
func main() {
	product := Product {"Kayak", "Water sport", 275}
	//var expense Expense = product
	var expense Expense = &product

	product.price = 100
	fmt.Println("Product field value:", product.price)
	fmt.Println("Expense method result:", expense.getCost(true))
}
*/