package main

type Product struct {
	Name, Category string
	Price float64
}

var Products = []Product {
	{"kayak", "watersports", 289},
	{"jacket", "watersports", 45.95},
	{"ball", "soccer", 18.95},
	{"stadium", "soccer", 18900},
	{"corner flags", "soccer", 57.85},
	{"thinking cap", "chess", 18},
	{"unsteady chair", "chess", 75},
	{"bling-bling king", "chess", 1200},
}