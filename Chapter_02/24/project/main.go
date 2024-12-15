package main
import (
	"net/http"
	"io"
)

type StringHandler struct {
	message string
}

func (sh StringHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	Printfln("Request for %v", request.URL.Path)
	switch request.URL.Path {
	case "/favicon.ico":
		http.NotFound(writer, request)
	case "/message":
		io.WriteString(writer, sh.message)
	default:
		http.Redirect(writer, request, "/message", http.StatusTemporaryRedirect)
	}
}

func main() {
	err := http.ListenAndServe(":5050", StringHandler{message: "Hello, World!!!!"})

	if (err != nil) {
		Printfln("Error: %v", err.Error())
	}
}