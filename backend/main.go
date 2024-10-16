package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// https://github.com/lib/pq/blob/master/README.md
// db orm?
// response format?

type Country struct {
	Id   int
	Name string
}

type City struct {
	Id         int
	Country_id int
	Name       string
}

type Citizen struct {
	Id      int
	City_id int
	Name    string
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
		w.WriteHeader(404)
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

func handleDatabase(w http.ResponseWriter, r *http.Request) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	sql := fmt.Sprintf("INSERT INTO countries(name) VALUES ('%s')", "moscow")
	db.Query(sql)

	rows, err := db.Query("SELECT * FROM countries")

	if err != nil {
		log.Fatal(err)
	}

	var country Country
	for rows.Next() {
		err := rows.Scan(&country.Id, &country.Name)
		if err != nil {
			log.Fatal(err)
		}
		result, err := json.Marshal(country)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(w, string(result))
	}

}

func main() {
	http.HandleFunc("GET /api/v1/country", getCountry)
	http.HandleFunc("GET /api/v1/countries", getCountries)
	http.HandleFunc("POST /api/v1/country", createCountry)
	http.HandleFunc("GET /api/v1/db", handleDatabase)

	http.ListenAndServe(":8080", nil)
}
