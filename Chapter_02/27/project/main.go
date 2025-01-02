package main

import (
	"reflect"
	"strings"
	//"fmt"
)

func incrementOrUpper(values ...interface{}) {
	for _, elem := range values {
		elemValue := reflect.ValueOf(elem)
		if (elemValue.CanSet()) {
			switch (elemValue.Kind()) {
			case reflect.Int:
				elemValue.SetInt(elemValue.Int() + 1)
			case reflect.String:
				elemValue.SetString(strings.ToUpper(elemValue.String()))
			}
			Printfln("Modigfied value: %v", elemValue)
		} else {
			Printfln("Cannot set %v : %v", elemValue.Kind(), elemValue)
		}
	}
}

func main() {
	name := "Alice"
	price := 255
	city := "London"

	incrementOrUpper(name, price, city)
	for _, val := range []interface{} {name, price, city} {
		Printfln("Value: %v", val)
	}
}