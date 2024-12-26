package main

func main() {
	listDrivers()
	db, err := openDatabase()
	if(err == nil) {
		db.Close()
	} else {
		panic(err)
	}
}