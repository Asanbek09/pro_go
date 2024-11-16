package main

import (
	"math/rand"
	"time"
)

/*
func IntRage(min, max int) int {
	return rand.Intn(max - min) + min
}
*/

var names = []string {"Alice", "Bob", "Charlie", "Dora", "Edith"}

func main() {

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(names), func (first, second int) {
		names[first], names[second] = names[second], names[first]
	})

	for i, name := range names {
		Printfln("Index %v: Name: %v", i, name)
	}

/*
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		Printfln("Value %v : %v", i, IntRage(10, 20))
	}
*/

/*
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		Printfln("Value %v : %v", i, rand.Intn(10))
		//Printfln("Value %v : %v", i, rand.Int())
	}
*/

/*
	val1 := 274.00
	val2 := 564.94

	Printfln("Abs: %v", math.Abs(val1))
	Printfln("Ceil: %v", math.Ceil(val2))
	Printfln("CopySign: %v", math.Copysign(val1, -5))
	Printfln("Floor: %v", math.Floor(val2))
	Printfln("Max: %v", math.Max(val1, val2))
	Printfln("Min: %v", math.Min(val1, val2))
	Printfln("Mod: %v", math.Mod(val1, val2))
	Printfln("Pow: %v", math.Pow(val1, val2))
	Printfln("Round: %v", math.Round(val2))
	Printfln("RoundToEleven: %v", math.RoundToEven(val2))
*/
}