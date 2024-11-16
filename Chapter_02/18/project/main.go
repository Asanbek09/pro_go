package main

import "math"

func main() {
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
}