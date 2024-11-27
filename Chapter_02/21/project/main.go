package main

import (
	"strings"
	"fmt"
	"encoding/json"
)

func main() {
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProduct {
		Product: &Kayak,
		Discount: 10.45,
	}
	encoder.Encode(&dp)
	fmt.Print(writer.String())

	/*
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	encoder.Encode(Kayak)
	fmt.Print(writer.String())
	*/

	/*
	m := map[string]float64 {
		"kayak": 259,
		"jacket": 58.75,
	}

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	encoder.Encode(m)
	fmt.Print(writer.String())
	*/
	
	/*
	names := []string {"Kayak", "jacket", "ball"}
	numbers := []int {10, 20, 30}
	var byteArray [5]byte
	copy(byteArray[0:], []byte(names[0]))
	byteSlice := []byte(names[0])

	var writer strings.Builder
	encoder := json.NewEncoder(&writer)

	encoder.Encode(names)
	encoder.Encode(numbers)
	encoder.Encode(byteArray)
	encoder.Encode(byteSlice)

	fmt.Print(writer.String())
	*/

	/*
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
	*/
}