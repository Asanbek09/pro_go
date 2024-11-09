package main

type Product struct {
	Name, Category string
	Price float64
}

var Kayak = Product {
	Name: "Kayak",
	Category: "water sports",
	Price: 255,
}

var Products = []Product {
	{"kayak", "water sports", 255},
	{"jacket", "water sports", 357.95},
	{"ball", "soccer", 17.55},
	{"flags", "soccer", 23.45},
	{"stadium", "soccer", 78950},
	{"cap", "chess", 65.75},
	{"chair", "chess", 657.40},
	{"bling king", "chess", 340},
}