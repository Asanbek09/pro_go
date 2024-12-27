package main

import "database/sql"

func queryDatabase(db *sql.DB) {
	rows, err := db.Query("Select * from Products")
	if (err == nil) {
		for (rows.Next()) {
			var id, category string
			var name string
			var price string
			rows.Scan(&id, &name, &category, &price)
			Printfln("Row: %v %v %v %v", id, name, category, price)
		}
	} else {
		Printfln("Error: %v", err)
	}
}

func main() {
	//listDrivers()
	db, err := openDatabase()
	if(err == nil) {
		queryDatabase(db)
		db.Close()
	} else {
		panic(err)
	}
}