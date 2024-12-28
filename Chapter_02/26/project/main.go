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

func insertRow(db *sql.DB, p *Product) (id int64) {
	res, err := db.Exec(`
	insert into products(name, category, price)
	values(?, ?, ?)`, p.Name, p.Category.Id, p.Price)
	if(err == nil) {
		id, err = res.LastInsertId()
		if (err != nil) {
			Printfln("Result Error: %v", err.Error())
		}
	} else {
		Printfln("Exec Error: %v", err.Error())
	}
	return
}

func main() {
	db, err := openDatabase()
	if (err == nil) {
		newProduct := Product {Name: "Statidum", Category: Category{Id:2}, Price: 79500}
		newId := insertRow(db, &newProduct)
		p := queryDatabase(db, int(newId))
		Printfln("New Product: %v", p)
		db.Close()
	} else {
		panic(err)
	}
}