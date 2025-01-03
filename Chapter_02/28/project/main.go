package main

import (
	"reflect"
	//"strings"
	//"fmt"
)

func getFieldValues(s interface{}) {
	structValue := reflect.ValueOf(s)
	if structValue.Kind() == reflect.Struct {
		for i := 0; i < structValue.NumField(); i++ {
			fieldType := structValue.Type().Field(i)
			fieldVal := structValue.Field(i)
			Printfln("Name: %v, Type: %v, Value: %v", fieldType.Name, fieldType.Type, fieldVal)
		}
	} else {
		Printfln("Not a struct")
	}
}

func main() {
	product := Product {Name: "Kayak", Category: "Water sports", Price: 277}
	customer := Customer {Name: "Alice", City: "Boston"}
	purchase := Purchase {Customer: customer, Product: product, Total: 290, taxRate: 20}

	getFieldValues(purchase)
}