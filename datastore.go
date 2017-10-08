package main

import (
	"context"
	"log"
	"net/url"

	"cloud.google.com/go/datastore"
)

// CloudstoreDB defines a Google Cloud Datastore Client
type CloudstoreDB struct {
	client *datastore.Client
}

// NewDSClient creates and returns a new client to Google Cloud Datastore
func NewDSClient(ctx context.Context, projectID string) (client *CloudstoreDB, err error) {
	dsClient, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Println("error connecting to Google Cloud Datastore: ", err)
		return
	}

	client = &CloudstoreDB{
		client: dsClient,
	}

	return
}

// CreateObject creates an object in the Google Cloud Datastore
func (db *CloudstoreDB) CreateObject(input url.Values) (err error) {

	return
}

// ReadObject retrieves all the objects for project from Google Cloud Datastore
func (db *CloudstoreDB) ReadObject() (output []map[string]interface{}, err error) {

	return
}
