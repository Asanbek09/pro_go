package main

import (
	"reflect"
	//"strings"
	//"fmt"
)

func executeFirstVoidMethods(s interface{}) {
	sVal := reflect.ValueOf(s)
	for i := 0; i < sVal.NumMethod(); i++ {
		method := sVal.Type().Method(i)
		if (method.Type.NumIn() == 1) {
			results := method.Func.Call([]reflect.Value{sVal})
			Printfln("Typ: %v, Method: %v, Results: %v", sVal.Type(), method.Name, results)
			break 
		} else {
			Printfln("Skipping method %v %v", method.Name, method.Type.NumIn())
		}
	}
}


func main() {
	executeFirstVoidMethods(&Product {Name: "kayak", Price: 268})
}