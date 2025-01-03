package main

import (
	"reflect"
	//"strings"
	//"fmt"
)

func invokeFunction(f interface{}, params ...interface{}) {
	paramVals := []reflect.Value {}
	for _, p := range params {
		paramVals = append(paramVals, reflect.ValueOf(p))
	}
	funcVal := reflect.ValueOf(f)
	if (funcVal.Kind() == reflect.Func) {
		results := funcVal.Call(paramVals)
		for i, r := range results {
			Printfln("Result #%v : %v", i, r)
		}
	}
}

func main() {
	names := []string {"alice", "bob", "charlie"}
	invokeFunction(Find, names, "London", "bob")
}