package main


func main() {
	for _, p := range Products {
		Printfln("Product: %v, category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
	}
}
