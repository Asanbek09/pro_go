package main

import (
	"reflect"
	//"strings"
	//"fmt"
)

func setFieldValue(s interface{}, newVals map[string]interface{}) {
	structValue := reflect.ValueOf(s)
	if (structValue.Kind() == reflect.Ptr && structValue.Elem().Kind() == reflect.Struct) {
		for name, newValue := range newVals {
			fieldVal := structValue.Elem().FieldByName(name)
			if (fieldVal.CanSet()) {
				fieldVal.Set(reflect.ValueOf(newValue))
			} else if (fieldVal.CanAddr()) {
				ptr := fieldVal.Addr()
				if (ptr.CanSet()) {
					ptr.Set(reflect.ValueOf(newValue))
				} else {
					Printfln("Cannot set field via pointer")
				}
			} else {
				Printfln("cannot set field")
			}
		}
	} else {
		Printfln("Not a pointer to a struct")
	}
}

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

	setFieldValue(&purchase, map[string]interface{} {
		"City": "London", "Category": "Boats", "Total": 215.20,
	})

	getFieldValues(purchase)
}