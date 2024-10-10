package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// db orm?
// response format?

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

func getCountry(w http.ResponseWriter, r *http.Request) {
	c := Country{1, "Russia"}
	resp, err := json.Marshal(c)
	if err != nil {
		fmt.Fprintln(w, "get error")
		return
	}
	fmt.Fprintln(w, string(resp))
}

func createCountry(w http.ResponseWriter, r *http.Request) {
	var p Country
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		fmt.Fprintln(w, "post error")
		return
	}
	resp, err := json.Marshal(p)
	fmt.Fprintln(w, string(resp))
}

func getCountries(w http.ResponseWriter, r *http.Request) {
	countries := [2]Country{{1, "Russia"}, {2, "Ukraine"}}
	resp, err := json.Marshal(countries)
	if err != nil {
		fmt.Fprintln(w, "error getAll")
		return
	}
	fmt.Fprintln(w, string(resp))
}

func main() {
	http.HandleFunc("GET /api/v1/country", getCountry)
	http.HandleFunc("GET /api/v1/countries", getCountries)
	http.HandleFunc("POST /api/v1/country", createCountry)

	http.ListenAndServe(":8080", nil)
}
