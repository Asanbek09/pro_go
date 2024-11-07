package main

import (
	"fmt"
	"01/composition/store"
)



func main() {
	products := map[string]store.ItemForSale {
		"Kayak": store.NewBoat("Kayak", 279, 1, false),
		"Ball": store.NewProduct("ball", "soccer", 19.45),
	}

	for key, p := range products {
		fmt.Println("Key", key, "Price:", p.Price(0.2))
	}

	/*
	products := map[string]*store.Product {
		"Kayak": store.NewBoat("kayak", 279, 1, false),
		"Ball": store.NewProduct("Ball", "Soccer", 12.65),
	}

	for _, p := range products {
		fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
	}
	*/

	/*
	kayak := store.NewProduct("kayak", "water sports", 269)
	type OfferBundle struct {
		*store.SpecialDeal
		*store.Product
	}

	bundle := OfferBundle {
		store.NewSpecialDeal("Weekend Special", kayak, 50),
		store.NewProduct("jacket", "clothes", 79.47),
	}

	fmt.Println("Price", bundle.Price(0))
	*/

	/*
	product := store.NewProduct("kayak", "water sport", 245)

	deal := store.NewSpecialDeal("Weeekend Special", product, 50)

	Name, price, Price := deal.GetDetails()
	fmt.Println("Name:", Name)
	fmt.Println("Price Field:", price)
	fmt.Println("Price method:", Price)
	*/

	/*
	rentals := []*store.RentalBoat {
		store.NewRentalBoat("Rubber Ring", 10, 1, false, false, "N/A", "N/A"),
		store.NewRentalBoat("Yacht", 50000, 5, true, true, "Bob", "Alice"),
		store.NewRentalBoat("Super Yacht", 100000, 15, true, true, "Dora", "Charlie"),
	}

	for _, r := range rentals {
		fmt.Println("Rental Boat:", r.Name, "Rental Price:", r.Price(0.2), "Captain:", r.Captain)
	}

	*/

	/*
	boats := []*store.Boat {
		store.NewBoat("kayak", 275, 1, false),
		store.NewBoat("canoe", 400, 3, false),
		store.NewBoat("tender", 650.25, 2, true),
	}

	for _, b := range boats {
		fmt.Println("Boat:", b.Name, "Price:", b.Price(0.2))
	}

	*/

	/*
	for _, b := range boats {
		fmt.Println("Convential:", b.Product.Name, "Direct:", b.Name)
	}
	*/

	/*
	kayak := store.NewProduct("kayak", "water sport", 275)
	jacket := &store.Product{Name: "jacklet", Category: "clothes"}

	for _, p := range []*store.Product {kayak, jacket} {
		fmt.Println("Name:", p.Name, "Category:", p.Category, "Price:", p.Price(0.2))
	}
		*/
}