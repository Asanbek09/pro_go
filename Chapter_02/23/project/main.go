package main

import (
	"html/template"
	"os"
	//"slices"
)

func GetCategories(products []Product) (categories []string) {
	catMap := map[string]string {}
	for _, p := range products {
		if (catMap[p.Category] == "") {
			catMap[p.Category] = p.Category
			categories = append(categories, p.Category)
		}
	}
	return
}

func Exec(t *template.Template) error {
	return t.Execute(os.Stdout, Products)
}

func main() {
	allTemplates := template.New("allTemplates")
	allTemplates.Funcs(map[string]interface{}{
		"getCats": GetCategories,
	})
	allTemplates, err := allTemplates.ParseGlob("templates/*.html")
	//allTemplates, err := template.ParseFiles("templates/template.html", "templates/list.html")
	if err == nil {
		selectedTemplated := allTemplates.Lookup("mainTemplate")
		err = Exec(selectedTemplated)
	}
	if err != nil {
		Printfln("Error: %v", err.Error())
	}
}

/*
allTemplates, err := template.ParseGlob("templates/*.html")
	if (err == nil) {
		for _, t := range allTemplates.Templates() {
			Printfln("Template name: %v", t.Name())
		}
	} else {
		Printfln("Error: %v", err.Error())
	}
*/

/*
allTemplates, err1 := template.ParseFiles("templates/template.html", "templates/extras.html")
	if (err1 == nil) {
		allTemplates.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
		os.Stdout.WriteString("\n")
		allTemplates.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
	} else {
		Printfln("Error: %v", err1.Error())
	}
*/

/*
t1, err1 := template.ParseFiles("templates/template.html")
	t2, err2 := template.ParseFiles("templates/extras.html")
	if (err1 == nil && err2 == nil) {
		t1.Execute(os.Stdout, &Kayak)
		os.Stdout.WriteString("\n")
		t2.Execute(os.Stdout, &Kayak)
	} else {
		Printfln("Err: %v %v", err1.Error(), err2.Error())
	}
*/

/*
t, err := template.ParseFiles("templates/template.html")
	if (err == nil) {
		t.Execute(os.Stdout, &Kayak)
	} else {
		Printfln("Error: %v", err.Error())
	}
*/
