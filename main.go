package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Hero struct {
	Name          string `json:"name"`
	Age           int    `json:"age"`
	RealFirstName string `json:"realFirstName"`
	RealLastName  string `json:"realLastName"`
}

var heroes = []Hero{
	{"superman", 30, "Clark", "Kent"},
	{"batman", 40, "Bruce", "Wayne"},
	{"flash", 25, "Barry", "Allen"},
}

func getHeroes(resp http.ResponseWriter, req *http.Request) {
	log.Printf("Got request to GET heroes\n")
	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(heroes)
}

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/heroes", getHeroes).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Listening on for requests, Port=%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
