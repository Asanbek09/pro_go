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

func insertAndUseCategory(db *sql.DB, name string, productIDs ...int) (err error) {
	tx, err := db.Begin()
	updatedFailed := false
	if (err == nil) {
		catResult, err := tx.Stmt(insertNewCategory).Exec(name)
		if (err == nil) {
			newID, _ := catResult.LastInsertId()
			preparedStatement := tx.Stmt(changeProductCategory)
			for _, id := range productIDs {
				changeResult, err := preparedStatement.Exec(newID, id)
				if (err == nil) {
					changes, _ := changeResult.RowsAffected()
					if (changes == 0) {
						updatedFailed = true
						break
					}
				}
			}
		}
	}
	if(err != nil || updatedFailed) {
		Printfln("Aborting transaction %v", err)
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return
}

func main() {
	db, err := openDatabase()
	if (err == nil) {
		insertAndUseCategory(db, "Category_1", 2)
		p := queryDatabase(db, 2)
		Printfln("Product: %v", p)
		insertAndUseCategory(db, "Category_2", 100)
		db.Close()
	} else {
		panic(err)
	}
}