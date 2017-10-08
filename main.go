package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"google.golang.org/appengine"
)

const keyKind = "flexstore"

func main() {

	var err error

	// check if GOOGLE_APPLICATION_CREDENTIALS is set in the local environment
	if appengine.IsDevAppServer() {
		appCred := os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
		if appCred == "" {
			log.Fatalf(`Error: environment variable GOOGLE_APPLICATION_CREDENTIALS is not set,
				please refer https://cloud.google.com/docs/authentication/getting-started 
				to setup authentication for your application`)
			os.Exit(-1)
		}
	}

	// retrieve the Google Project ID from environment
	projID := os.Getenv("FLEXSTORE_PROJECT_ID")
	if projID == "" {
		log.Fatal(`Error: environment variable FLEXSTORE_PROJECT_ID is not set,
			please set this variable to your Google Cloud Project.`)
		os.Exit(-1)
	}

	ctx := context.Background()

	// create a connection to Google Cloud Datastore
	dsClient, err := NewDSClient(ctx, projID)
	if err != nil {
		os.Exit(-1)
	}

	// setup the HTTP routes and handlers
	r := mux.NewRouter().StrictSlash(false)
	r.HandleFunc("/", DefaultHandler)
	r.HandleFunc("/liveness_check", LivenessCheckHandler)

	r.HandleFunc("/save", func(w http.ResponseWriter, r *http.Request) {
		PostValHandler(ctx, w, r, dsClient)
	}).Methods("GET")

	r.HandleFunc("/retrieve", func(w http.ResponseWriter, r *http.Request) {
		GetValHandler(ctx, w, r, dsClient)
	}).Methods("GET")

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Printf("Flexstore is now serving RESTfully [port %s]...", server.Addr)
	if err = server.ListenAndServe(); err != nil {
		log.Println(err)
		dsClient.client.Close()
	}
}
