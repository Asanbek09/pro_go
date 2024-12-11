package main

import (
	"html/template"
	"os"
)

func main() {
	allTemplates, err1 := template.ParseFiles("templates/template.html", "templates/extras.html")
	if (err1 == nil) {
		allTemplates.ExecuteTemplate(os.Stdout, "template.html", &Kayak)
		os.Stdout.WriteString("\n")
		allTemplates.ExecuteTemplate(os.Stdout, "extras.html", &Kayak)
	} else {
		Printfln("Error: %v", err1.Error())
	}
}

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