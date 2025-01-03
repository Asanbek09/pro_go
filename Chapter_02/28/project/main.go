package main

import (
	"reflect"
	//"strings"
	//"fmt"
)

func setMap(m interface{}, key interface{}, val interface{}) {
	mapValue := reflect.ValueOf(m)
	keyValue := reflect.ValueOf(key)
	valValue := reflect.ValueOf(val)
	if (mapValue.Kind() == reflect.Map && mapValue.Type().Key() == keyValue.Type() && mapValue.Type().Elem() == valValue.Type()) {
		mapValue.SetMapIndex(keyValue, valValue)
	} else {
		Printfln("Not a map or mismatched types")
	}
}

func removeFromMap(m interface{}, key interface{}) {
	mapValue := reflect.ValueOf(m)
	keyValue := reflect.ValueOf(key)
	if (mapValue.Kind() == reflect.Map && mapValue.Type().Key() == keyValue.Type()) {
		mapValue.SetMapIndex(keyValue, reflect.Value{})
	}
}

func main() {
	pricesMap := map[string]float64 {
		"kayak": 299, "jacket": 45.98, "soccer ball": 15.69,
	}
	setMap(pricesMap, "Kayak", 56.00)
	setMap(pricesMap, "Hat", 25.00)
	removeFromMap(pricesMap, "jacket")
	for k, v := range pricesMap {
		Printfln("Key: %v, Value: %v", k, v)
	}
}