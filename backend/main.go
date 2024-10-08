package main

import (
	"fmt"
	"html"
	"net/http"
)

func handleFnc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

}

func main() {
	http.HandleFunc("/api/v1/country", handleFnc)

	http.ListenAndServe(":8080", nil)
}
