package main

import (
	"net/http"
	"os"
	"time"
	"io"
	"encoding/json"
	"strings"
	//"net/url"
)

func main() {
	go http.ListenAndServe(":5550", nil)
	time.Sleep(time.Second)

	var builder strings.Builder
	err := json.NewEncoder(&builder).Encode(Products[0])
	if (err == nil) {
		req, err := http.NewRequest(http.MethodPost, "http://localhost:5550/echo", io.NopCloser(strings.NewReader(builder.String())))
		if (err == nil) {
			req.Header["Content-Type"] = []string{"application/json"}
			response, err := http.DefaultClient.Do(req)
			if (err == nil && response.StatusCode == http.StatusOK) {
				io.Copy(os.Stdout, response.Body)
				defer response.Body.Close()
			} else {
				Printfln("Request Error: %v", err.Error())
			}
		} else {
			Printfln("Request Init Error: %v", err.Error())
		}
	} else {
		Printfln("Encoder Error: %v", err.Error())
	}
}