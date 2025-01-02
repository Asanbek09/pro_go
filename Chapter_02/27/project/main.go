package main

import (
	"reflect"
	//"strings"
	//"fmt"
)
var intPtrType = reflect.TypeOf((*int)(nil))
var byteSliceType = reflect.TypeOf([]byte(nil))

func printDetails(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		elemType := reflect.TypeOf(elem)
		if (elemType == intPtrType) {
			Printfln("Pointer to Int: %v", elemValue.Elem().Int())
		} else if (elemType == byteSliceType){
			Printfln("Byte slice: %v", elemValue.Bytes())
		} else {
			switch elemValue.Kind() {
			case reflect.Bool:
				var val bool = elemValue.Bool()
				Printfln("Bool: %v", val)
			case reflect.Int:
				var val int64 = elemValue.Int()
				Printfln("Int: %v", val)
			case reflect.Float32, reflect.Float64:
				var val float64 =  elemValue.Float()
				Printfln("Float: %v", val)
			case reflect.String:
				var val string = elemValue.String()
				Printfln("String: %v", val)
			//case reflect.Ptr:
			//	var val reflect.Value = elemValue.Elem()
			//	if (val.Kind() == reflect.Int) {
			//		Printfln("Pointer to Int: %v", val.Int())
			//	}
			default:
				Printfln("Other: %v", elemValue.String())
			}
		}
	}
}

func main() {
	product := Product {
		Name: "kayak", Category: "water sports", Price: 289,
	}
	number := 100
	slice := []byte("Alice")
	printDetails(true, 10, 20.50, "Alice", &number, product, slice)
}