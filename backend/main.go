package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// now
// github create repo
// single branch master
// 5 CRUD endpoints country + entity. Do they have entities in vainla go?
// 5 CRUD endpoints city + entity
// 5 CRUD endpoints citizen + entity
// dummy responses.
// relations?
// connect with database, postgresql?
// elasticsearch? ES doesn't have native web interface.
// simple frontend NEXT.js
// end

type Country struct {
	Id   int
	Name string
}

type city struct {
	id         int
	country_id int
	name       string
}

type citizen struct {
	id      int
	city_id int
	name    string
}

func handleFnc(w http.ResponseWriter, r *http.Request) {
	c := Country{1, "Russia"}
	// fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	resp, err := json.Marshal(c)
	if err != nil {
		fmt.Fprintln(w, "error")
	}
	fmt.Fprintln(w, string(resp))

}

func main() {
	http.HandleFunc("/api/v1/country", handleFnc)

	http.ListenAndServe(":8080", nil)
}
