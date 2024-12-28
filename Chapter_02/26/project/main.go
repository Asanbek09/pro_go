package main

import (
	"database/sql"
	//"go/printer"
)

type Category struct {
	Id int
	Name string
}

type Product struct {
	Id int
	Name string
	Category
	Price float64
}

func queryDatabase(db *sql.DB) (products []Product, err error) {
	rows, err := db.Query(`select products.id, products.name, products.price, categories.id as "category.id", categories.name as "category.name" 
	from products, categories where products.category = categories.id`)
	if (err != nil) {
		return
	} else {
		results, err := scanIntoStruct(rows, &Product{})
		if err == nil {
			products = (results).([]Product)
		} else {
			Printfln("Scanning error: %v", err)
		}
	}
	return
} 


func main() {
	db, err := openDatabase()
	if (err == nil) {
		products, _ := queryDatabase(db)
		for _, p := range products {
			Printfln("Product: %v", p)
		}
		db.Close()
	} else {
		panic(err)
	}
}