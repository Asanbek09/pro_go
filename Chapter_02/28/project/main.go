package main

import (
	"reflect"
	//"strings"
	//"fmt"
)

func printMapContents(m interface{}) {
	mapValue := reflect.ValueOf(m)
	if (mapValue.Kind() == reflect.Map) {
		iter := mapValue.MapRange()
		for iter.Next() {
			Printfln("Map Key: %v, Value: %v", iter.Key(), iter.Value())
		}
	} else {
		Printfln("Not a map")
	}
}

func main() {
	pricesMap := map[string]float64 {
		"kayak": 299, "jacket": 45.98, "soccer ball": 15.69,
	}
	printMapContents(pricesMap)
}