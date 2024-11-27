package main

import (
	"strings"
	"fmt"
	"encoding/json"
)

func main() {
	var b bool = true
	var str string = "Hello"
	var fval float64 = 99.98
	var ival int = 100
	var pointer *int = &ival

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	for _, val := range []interface{} {b, str, fval, ival, pointer} {
		encoder.Encode(val)
	}

	fmt.Print(writer.String())
}