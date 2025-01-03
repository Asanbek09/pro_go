package main

import (
	"reflect"
	//"strings"
	//"fmt"
)

func executeFirstVoidMethods(s interface{}) {
	sVal := reflect.ValueOf(s)
	for i := 0; i < sVal.NumMethod(); i++ {
		method := sVal.Method(i)
		if (method.Type().NumIn() == 0) {
			results := method.Call([]reflect.Value{})
			Printfln("Type: %v, Method: %v, Results: %v", sVal.Type(), sVal.Type().Method(i).Name, results)
			break
		} else {
			Printfln("Skipping method %v %v", sVal.Type().Method(i).Name, method.Type().NumIn())
		}
	}
}


func main() {
	executeFirstVoidMethods(&Product {Name: "kayak", Price: 268})
}