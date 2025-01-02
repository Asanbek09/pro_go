package main

import (
	"reflect"
	//"strings"
	//"fmt"
)

func describeMap(m interface{}) {
	mapType := reflect.TypeOf(m)
	if (mapType.Kind() == reflect.Map) {
		Printfln("Key type: %v, Val type: %v", mapType.Key(), mapType.Elem())
	} else {
		Printfln("Not a map")
	}
}

func main() {
	pricesMap := map[string]float64 {
		"kayak": 299, "jacket": 45.98, "soccer ball": 15.69,
	}
	describeMap(pricesMap)
}