package main

import "strconv"

type Product struct {
	Name, Category string
	Price float64
}

type ProductSlice []*Product

var Products = ProductSlice {
	{"kayak", "water sport", 278},
	{"jacket", "water sport", 45.95},
	{"ball", "soccer", 14.85},
	{"flags", "soccer", 17.90},
	{"stadium", "soccer", 89650},
	{"cap", "chess", 17},
	{"chair", "chess", 27.85},
	{"bling king", "chess", 1200},
}

func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}