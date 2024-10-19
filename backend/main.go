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
	var country Country
	countryId := r.PathValue("country")
	query := fmt.Sprintf("SELECT * FROM countries WHERE id = %s", countryId)
	rows, err := db.Query(query)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintln(w, err)
		return
	}
	rows.Next()
	err = rows.Scan(&country.Id, &country.Name)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintln(w, err)
		return
	}
	json, err := json.Marshal(country)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintln(w, "Country stringify error")
		return
	}
	fmt.Fprintln(w, string(json))
}

func createCountry(w http.ResponseWriter, r *http.Request) {

	var p Country
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintln(w, "body error")
		return
	}
	sql := fmt.Sprintf("INSERT INTO countries(name) VALUES ('%s')", p.Name)
	if _, err := db.Query(sql); err != nil {
		w.WriteHeader(404)
		fmt.Fprintln(w, "db insert error")
		return
	}
}

func getCountries(w http.ResponseWriter, r *http.Request) {
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

var db *sql.DB

func initializeDatabaseFromEnv() {
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
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
}

func runEndpoints() {
	http.HandleFunc("GET /api/v1/country/{country}", getCountry)
	http.HandleFunc("GET /api/v1/countries", getCountries)
	http.HandleFunc("POST /api/v1/country", createCountry)

	http.ListenAndServe(":8080", nil)
}

func main() {
	initializeDatabaseFromEnv()
	runEndpoints()
}
