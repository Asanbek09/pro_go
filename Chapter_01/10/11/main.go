package main

import "fmt"

func main() {
	type Product struct {
		name, category string
		price float64
	}

	type StockLevel struct {
		Product
		Alternate Product
		count int
	}

	array := [1]StockLevel {
		{
			Product: Product{"kayak", "water sport", 255.00},
			Alternate: Product{"jacket", "soccre", 150.95},
			count: 50,
		},
	}
	fmt.Println("array:", array[0].Product.name)

	slice := []StockLevel {
		{
			Product: Product{"kayak", "water sport", 255.00},
			Alternate: Product{"jacket", "clothes", 149.6},
			count: 25,
		},
	}
	fmt.Println("Slice:", slice[0].Product.name)

	kvp := map[string]StockLevel {
		"kayak": {
			Product: Product{"kayak", "water sport", 250.00},
			Alternate: Product{"jacket", "clothes", 120.54},
		},		
	}
	fmt.Println("Map:", kvp["kayak"].Product.name)
}