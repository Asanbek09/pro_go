package main

import (
	"encoding/json"
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader(`true "Hello" 99.99 200`)
	vals := []interface{} { }

	decoder := json.NewDecoder(reader)

	for {
		var decodedVal interface{}
		err := decoder.Decode(&decodedVal)
		if (err != nil) {
			if(err != io.EOF) {
				Printfln("Error: %v", err.Error())
			}
			break
		}
		vals = append(vals, decodedVal)
	}

	for _, val := range vals {
		if num, ok := val.(json.Number); ok {
			if ival, err := num.Int64(); err == nil {
				Printfln("Decoded Integer: %v", ival)
			} else if fpval, err := num.Float64(); err == nil {
				Printfln("Decoded floating point: %v", fpval)
			} else {
				Printfln("Decoded String: %v", num.String())
			}
		} else {
			Printfln("Decoded (%T): %v", val, val)
		}
	}


	/*
	var writer strings.Builder
	encoder := json.NewEncoder(&writer)
	dp := DiscountedProduct {
		Product: &Kayak,
		Discount: 10.45,
	}
	
	namedItems := []Named {&dp, &Person{PersonName: "Alice"}}
	encoder.Encode(namedItems)
	
	//fmt.Print(writer.String())
	*/

	//encoder.Encode(&dp)
	//dp2 := DiscountedProduct {Discount: 10.45}
	//encoder.Encode(&dp2)

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