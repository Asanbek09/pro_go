package main

import (
	//"math/rand"
	//"time"
	//"sort"
)

func main() {
	products := []Product {
		{"kayak", 245},
		{"jacket", 273.95},
		{"ball", 15.75},
	}
	ProductSlicesByName(products)

	SortWith(products, func (p1, p2 Product) bool {
		return p1.Name < p2.Name
	})

	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}

/*
	ProductSlices(products)

	for _, p := range products {
		Printfln("Name: %v, Price: %.2f", p.Name, p.Price)
	}

*/

/*	
	ints := []int {9, 4, 5, -1, -4, 2}
	
	sortedInts := make([]int, len(ints))
	copy(sortedInts, ints)
	sort.Ints(sortedInts)
	Printfln("Ints: %v", ints)
	Printfln("Ints sorted: %v", sortedInts)

	index0f4 := sort.SearchInts(sortedInts, 4)
	index0f3 := sort.SearchInts(sortedInts, 3)
	Printfln("Index of 4: %v (present: %v)", index0f4, sortedInts[index0f4] == 4)
	Printfln("Index of 3: %v (present: %v)", index0f3, sortedInts[index0f3] == 3)
*/
	
	//Printfln("Index of 4: %v", index0f4)
	//Printfln("Index of 3: %v", index0f3)
 
	/*
	ints := []int {9, 3, 1, 9, 4, -1, -5, 7, 4}
	Printfln("Ints: %v", ints)
	sort.Ints(ints)
	Printfln("Ints sorted: %v", ints)

	floats := []float64 {254, 289.94, 14.50, 23.80, 26, 27.95}
	Printfln("Floats: %v", floats)
	sort.Float64s(floats)
	Printfln("Floats sorted: %v", floats)

	strings := []string {"ball", "cap", "stadium", "jacket", "map"}
	Printfln("Strings: %v", strings)
	if (!sort.StringsAreSorted(strings)) {
		sort.Strings(strings)
		Printfln("Strings sorted: %v", strings)
	} else {
		Printfln("Strings already sorted: %v", strings)
	}
	*/
}