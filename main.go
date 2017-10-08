package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// PostValHandler save the values of 'input' param in google cloud datastore
func PostValHandler(w http.ResponseWriter, r *http.Request) {

}

// GetValHandler retrieve/return all the values saved so far
func GetValHandler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/save", PostValHandler).Methods("GET")
	r.HandleFunc("/retrieve", GetValHandler).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Println("Flexstore is now serving RESTfully...")
	server.ListenAndServe()
}
