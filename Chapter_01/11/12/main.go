package main

import "fmt"

type Expense interface {
	getName() string
	getCost(annual bool) float64
}

func calcTotal(expenses []Expense) (total float64) {
	for _, item := range expenses {
		total += item.getCost(true)
	}
	return
}

type Account struct {
	accountNumber int
	expenses []Expense
}

func main() {
	account := Account {
		accountNumber: 125476,
		expenses: []Expense {
			Product {"Kayak", "Water Sport", 255},
			Service {"Boat Cover", 19, 55.65},
		},
	}

	for _, expense := range account.expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}
	fmt.Println("Total:", calcTotal(account.expenses))
}


/*
func main() {
	
	kayak := Product {"Kayak", "water sport", 255}
	insurance := Service {"Boat Cover", 18, 99.55}

	fmt.Println("Product:", kayak.name, "Price:", kayak.price)
	fmt.Println("Service:", insurance.description, "Price:", insurance.monthlyFee * float64(insurance.durationMonths))
	

	expenses := []Expense {
		Product {"Kayak", "Water Sport", 255},
		Service {"Boat Cover", 12, 89.65},
	}

	for _, expense := range expenses {
		fmt.Println("Expense:", expense.getName(), "Cost:", expense.getCost(true))
	}

	fmt.Println("Total:", calcTotal(expenses))
}
*/