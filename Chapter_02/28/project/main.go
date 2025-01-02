package main

import (
	"reflect"
	//"strings"
	//"fmt"
)

func setValue(arrayOrSlice interface{}, index int, replacement interface{}) {
	arrayOrSliceVal := reflect.ValueOf(arrayOrSlice)
	replacementVal := reflect.ValueOf(replacement)
	if (arrayOrSliceVal.Kind() == reflect.Slice) {
		elemVal := arrayOrSliceVal.Index(index)
		if (elemVal.CanSet()) {
			elemVal.Set(replacementVal)
		}
	} else if (arrayOrSliceVal.Kind() == reflect.Ptr && arrayOrSliceVal.Elem().Kind() == reflect.Array && arrayOrSliceVal.Elem().CanSet()) {
		arrayOrSliceVal.Elem().Index(index).Set(replacementVal)
	}
}

func main() {
	name := "Alice"
	city := "London"
	hobby := "Running"

	slice := []string {name, city, hobby}
	array := [3]string {name, city, hobby}

	Printfln("Original Slice: %v", slice)
	newCity := "Paris"
	setValue(slice, 1, newCity)
	Printfln("Modified slice: %v", slice)

	Printfln("Original slice: %v", array)
	newCity = "Rome"
	setValue(&array, 1, newCity)
	Printfln("Modified slice: %v", array)
}