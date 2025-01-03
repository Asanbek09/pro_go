package main

import (
	"reflect"
	//"strings"
	//"fmt"
)

func printMapContents(m interface{}) {
	mapValue := reflect.ValueOf(m)
	if (mapValue.Kind() == reflect.Map) {
		for _, keyVal := range mapValue.MapKeys() {
			reflectedVal := mapValue.MapIndex(keyVal)
			Printfln("Map Key: %v, Value: %v", keyVal, reflectedVal)
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