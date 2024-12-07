package main

import (
	//"fmt"
	//"time"
	//"encoding/json"
	"os"
	//"path/filepath"
)

func main() {
	path, err := os.Getwd()
	if (err == nil) {
		dirEntries, err := os.ReadDir(path)
		if (err == nil) {
			for _, dentry := range dirEntries {
				Printfln("Entry name: %v, IsDir: %v", dentry.Name(), dentry.IsDir())
			}
		}
	}
	if (err != nil) {
		Printfln("Error: %v", err.Error())
	}
}
/*
path, err := os.UserHomeDir()
	if err == nil {
		path = filepath.Join(path, "MyApp", "MyTempFile.json")
	}

	Printfln("Full path: %v", path)

	err = os.MkdirAll(filepath.Dir(path), 0766)
	if (err == nil) {
		file, err := os.OpenFile(path, os.O_CREATE | os.O_WRONLY, 0666)
		if(err == nil) {
			defer file.Close()
			encoder := json.NewEncoder(file)
			encoder.Encode(Products)
		}
	}
	if(err != nil) {
		Printfln("Error: %v", err.Error())
	}
*/

//Printfln("Volume name: %v", filepath.VolumeName(path))
//Printfln("Dir component: %v", filepath.Dir(path))
//Printfln("File component: %v", filepath.Base(path))
//Printfln("File extension: %v", filepath.Ext(path))

/*
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
*/

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
