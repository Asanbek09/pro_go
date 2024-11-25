package main

type Product struct {
	Name, Category string
	Price float64
}

var Kayak = Product {
	Name: "kayak",
	Category: "water sports",
	Price: 274.90,
}

var Products = []Product {
	{"kayak", "water sports", 264.55},
	{"jacket", "leather", 283},
	{"ball", "soccer", 18.95},
	{"flags", "soccer", 34.94},
	{"stadium", "soccer", 198000},
	{"thinking cap", "chess", 15},
	{"unsteady chair", "chess", 29.95},
	{"bling-bling king", "chess", 1200},
}