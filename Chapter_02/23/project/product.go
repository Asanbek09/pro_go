package main

type Product struct {
	Name, Category string
	Price float64
}

var Kayak = Product {
	Name: "kayak",
	Category: "water sports",
	Price: 259,
}

var Products = []Product {
	{"kayak", "water sport", 243},
	{"jacket", "water sport", 48.55},
	{"ball", "soccer", 17.95},
	{"flags", "soccer", 47.50},
	{"stadium", "soccer", 98500},
	{"cap", "chess", 166},
	{"chair", "chess", 75},
	{"bling-bling king", "chess", 1105},
}

func (p *Product) AddTax() float64 {
	return p.Price * 1.2
}

func (p *Product) ApplyDiscount(amount float64) float64 {
	return p.Price - amount
}