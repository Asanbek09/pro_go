package main

import "strconv"

type Product struct {
	Name, Category string
	Price float64
}

var Productlist =[]*Product {
	{"kayak", "water sports", 256},
	{"jacket", "clothes", 49.56},
	{"ball", "soccer", 25.65},
	{"corner flags", "soccer", 34.95},
	{"stadium", "athletic", 89500},
	{"cap", "chess", 16},
	{"chair", "chess", 75},
	{"bling king", "chess", 1200},
}

type ProductGroup []*Product
type ProductData = map[string]ProductGroup

var Products = make(ProductData)
func ToCurrency(val float64) string {
	return "$" + strconv.FormatFloat(val, 'f', 2, 64)
}

func init() {
	for _, p := range Productlist {
		if _, ok := Products[p.Category]; ok {
			Products[p.Category] = append(Products[p.Category], p)
		} else {
			Products[p.Category] = ProductGroup{ p }
		}
	}
}