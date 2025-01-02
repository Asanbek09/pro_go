package main

import (
	"reflect"
	//"strings"
	//"fmt"
)

func getTypePath(t reflect.Type) (path string) {
	path = t.PkgPath()
	if (path == "") {
		path = "(built-in)"
	}
	return
}

func printDetails(values ...interface{}) {
	for _, elem := range values {
		elemType := reflect.TypeOf(elem)
		Printfln("Name: %v, PkgPath: %v, Kind: %v", elemType.Name(), getTypePath(elemType), elemType.Kind())
	}
}

type Payment struct {
	Currency string
	Amount float64
}

func main() {
	product := Product {
		Name: "kayak", Category: "water sports", Price: 289,
	}
	customer := Customer {Name: "alice", City: "Boston"}
	payment := Payment {Currency: "USD", Amount: 100.45}
	printDetails(product, customer, payment, 10, true)
}