package main

import (
	//"fmt"
	//"time"
	"encoding/json"
	"os"
)

func main() {
	cheapProducts := []Product{}
	for _, p := range Products {
		if p.Price < 100 {
			cheapProducts = append(cheapProducts, p)
		}
	}

	file, err := os.CreateTemp(".", "tempfile-*.json")
	if err == nil {
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err.Error())
	}
}

/*
cheapProducts := []Product {}
	for _, p := range Products {
		if (p.Price < 100) {
			cheapProducts = append(cheapProducts, p)
		}
	}

	file, err := os.OpenFile("cheap.json", os.O_WRONLY | os.O_CREATE, 0666)
	if (err == nil) {
		defer file.Close()
		encoder := json.NewEncoder(file)
		encoder.Encode(cheapProducts)
	} else {
		Printfln("Error: %v", err.Error())
	}
*/

/*
total := 0.0
	for _, p := range Products {
		total += p.Price
	}

	dataStr := fmt.Sprintf("Time: %v, Total: $%.2f\n", time.Now().Format("Mon 14:05:09"), total)

	file, err := os.OpenFile("output.txt", os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0666)

	if (err == nil) {
		defer file.Close()
		file.WriteString(dataStr)
	} else {
		Printfln("Error: %v", err.Error())
	}
*/

/*
total := 0.0
	for _, p := range Products {
		total += p.Price
	}

	dataStr := fmt.Sprintf("Time: %v, Total: $%.2f\n", time.Now().Format("Mon 15:05:03"), total)

	err := os.WriteFile("output.txt", []byte(dataStr), 0666)
	if (err == nil) {
		fmt.Println("Output file created")
	} else {
		Printfln("Error: %v", err.Error())
	}
*/

/*
for _, p := range Products {
		Printfln("Product: %v, category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
	}
*/
