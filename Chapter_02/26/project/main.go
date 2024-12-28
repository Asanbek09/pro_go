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

func queryDatabase(db *sql.DB, id int) (p Product) {
	row := db.QueryRow(`select products.id, products.name, products.price, categories.id as cat_id, 
	categories.name as catname from products, categories 
	where products.category = categories.id and products.id = ?`, id)
	if(row.Err() == nil) {
		scanErr := row.Scan(&p.Id, &p.Name, &p.Price, &p.Category.Id, &p.Category.Name)
		if (scanErr != nil) {
			Printfln("Scan error: %v", scanErr)
		}
	} else {
		Printfln("Row error: %v", row.Err().Error())
	}
	return
}

func insertAndUseCategory(name string, productIDs ...int) {
	result, err := insertNewCategory.Exec(name)
	if (err == nil) {
		newID, _ := result.LastInsertId()
		for _, id := range productIDs {
			changeProductCategory.Exec(int(newID), id)
		}
	} else {
		Printfln("Prepared statement error: %v", err)
	}
}

func main() {
	db, err := openDatabase()
	if (err == nil) {
		insertAndUseCategory("Misc Products", 2)
		p := queryDatabase(db, 2)
		Printfln("Product: %v", p)
		db.Close()
	} else {
		panic(err)
	}
}