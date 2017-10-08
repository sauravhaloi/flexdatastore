package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// DB is an interface for CRUD operations of flexstore
type DB interface {
	CreateObject(obj []string) (err error)
	ReadObject() (obj []string, err error)
}

// DefaultHandler is the root route handler
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Fprint(w, "Hello world!")
}

// LivenessCheckHandler implements an application liveness checker for Google App Engine Flex apps
func LivenessCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	status := make(map[string]interface{})
	status["status"] = "Service is Up and Running"

	j, err := json.Marshal(status)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

// PostValHandler save the values of 'input' param in google cloud datastore
func PostValHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, db *CloudstoreDB) {
	params := r.URL.Query()["input"]

	if len(params) > 0 {
		err := db.CreateObject(params)
		if err != nil {
			log.Println(err)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
		}
	}

}

// GetValHandler retrieve/return all the values saved so far
func GetValHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, db *CloudstoreDB) {
	objs, err := db.ReadObject()
	if err != nil {
		log.Println(err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
	} else {
		j, err := json.Marshal(objs)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
