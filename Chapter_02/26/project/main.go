package main

import "database/sql"

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

func queryDatabase(db *sql.DB, categoryName string) []Product {
	products := []Product {}
	rows, err := db.Query(`select products.id, products.name, products.price, categories.id as cat_id, categories.name as catname from products, categories where products.category = categories.id and categories.name=?`, categoryName)
	if(err == nil) {
		for (rows.Next()) {
			p := Product{}
			scanErr := rows.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
			if(scanErr == nil) {
				products = append(products, p)
			} else {
				Printfln("Scan error: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error: %v", err)
	}
	return products
}

func main() {
	//listDrivers()
	db, err := openDatabase()
	if(err == nil) {
		for _, cat := range []string {"Soccer", "Watersports"} {
			Printfln("-----%v Results -----", cat)
			products := queryDatabase(db, cat)
			for i, p := range products {
				Printfln("#%v: %v %v %v", i, p.Name, p.Category, p.Category.Name, p.Price)
			}
		}
		db.Close()
	} else {
		panic(err)
	}
}