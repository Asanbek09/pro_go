package main

import (
	"net/http"
	"os"
	"time"
	"io"
	//"encoding/json"
)

func main() {
	go http.ListenAndServe(":5550", nil)
	time.Sleep(time.Second)

	formData := map[string][]string {
		"name": { "kayak" },
		"category": { "water sports" },
		"price": { "259" },
	}

	response, err := http.PostForm("http://localhost:5550/echo", formData)

	if(err == nil && response.StatusCode == http.StatusOK) {
		io.Copy(os.Stdout, response.Body)
		defer response.Body.Close()
	} else {
		Printfln("Error: %v, Status Code: %v", err.Error(), response.StatusCode)
	}
}