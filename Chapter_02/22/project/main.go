package main

import (
	"fmt"
	"time"
	"os"
)

func main() {
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
}


/*
for _, p := range Products {
		Printfln("Product: %v, category: %v, Price: $%.2f", p.Name, p.Category, p.Price)
	}
*/