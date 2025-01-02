package main

import (
	"reflect"
	"strings"
	"fmt"
)



func printDetails(values ...interface{}) {
	for _, elem := range values {
		fieldDetails := []string {}
		elemType := reflect.TypeOf(elem)
		elemValue := reflect.ValueOf(elem)
		if elemType.Kind() == reflect.Struct {
			for i := 0; i< elemType.NumField(); i++ {
				fieldName := elemType.Field(i).Name
				fieldVal := elemValue.Field(i)
				fieldDetails = append(fieldDetails, fmt.Sprintf("%v : %v", fieldName, fieldVal))
			}
			Printfln("%v : %v", elemType.Name(), strings.Join(fieldDetails, ", "))
		} else {
			Printfln("%v : %v", elemType.Name(), elemValue)
		}
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