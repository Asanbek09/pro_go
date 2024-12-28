package main

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

func listDrivers() {
	for _, driver := range sql.Drivers() {
		Printfln("Driver: %v", driver)
	}
}

var insertNewCategory *sql.Stmt
var changeProductCategory *sql.Stmt

func openDatabase() (db *sql.DB, err error) {
	db, err = sql.Open("sqlite", "products.db")
	if(err == nil) {
		Printfln("Opened database")
		insertNewCategory, _ = db.Prepare("insert into categories (name) values(?)")
		changeProductCategory, _ = db.Prepare("update products set category = ? where id = ?")
	}
	return
}