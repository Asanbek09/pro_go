package main

func printDetails(values ...interface{}) {
	for _, elem := range values {
		switch val := elem.(type) {
		case Product:
			Printfln("Product - Name: %v, Cateory: %v, Price: %v", val.Name, val.Category, val.Price)
		case Customer:
			Printfln("Customer - Name: %v, City: %v", val.Name, val.City)
		}
	}
}

func main() {
	product := Product {
		Name: "kayak", Category: "water sports", Price: 289,
	}
	customer := Customer {Name: "alice", City: "Boston"}
	printDetails(product, customer)
}