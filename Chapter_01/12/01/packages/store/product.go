package store

type Product struct {
	Name, Category string // имена полей с заглавной буквы означает что они экспортируются
	price float64 // имена полей с строчной буквы означает что доступ к нему возможен только внутри пакета store 
}