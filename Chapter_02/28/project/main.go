package main

import (
	"reflect"
	//"strings"
	//"fmt"
)

func inspectTags(s interface{}, tagName string) {
	structType := reflect.TypeOf(s)
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		tag := field.Tag
		valGet := tag.Get(tagName)
		valLookup, ok := tag.Lookup(tagName)
		Printfln("Field: %v, Tag %v : %v", field.Name, tagName, valGet)
		Printfln("Field: %v, Tag %v : %v, Set: %v", field.Name, tagName, valLookup, ok)
	}
}

type Person struct {
	Name string `alias:"id"`
	City string `alias:""`
	Country string
}



func main() {
	inspectTags(Person{}, "alias")
}