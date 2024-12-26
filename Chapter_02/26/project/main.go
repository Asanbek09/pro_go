package main

func main() {
	db, err := openDatabase()
	if(err == nil) {
		db.Close()
	} else {
		panic(err)
	}
}