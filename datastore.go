package main

import (
	"context"
	"log"
	"strings"

	"cloud.google.com/go/datastore"
)

// CloudstoreDB defines a Google Cloud Datastore Client
type CloudstoreDB struct {
	client *datastore.Client
}

// Entity represents the data to be stored/retrieved from Google Cloud Datastore
type Entity struct {
	Value string
}

// NewDSClient creates and returns a new client to Google Cloud Datastore
func NewDSClient(ctx context.Context, projectID string) (client *CloudstoreDB, err error) {
	ctx = context.Background()
	dsClient, err := datastore.NewClient(ctx, projectID)
	if err != nil {
		log.Println("error connecting to Google Cloud Datastore: ", err)
		return
	}
	log.Println("Successfully connected to Google Cloud Datastore for project " + projectID)

	client = &CloudstoreDB{
		client: dsClient,
	}

	return
}

// Close terminates an existing connection to Google Cloud Datastore
func Close(client *CloudstoreDB) {
	var err error
	err = client.client.Close()
	if err != nil {
		log.Println(err)
	}

}

// CreateObject creates an object in the Google Cloud Datastore
func (db *CloudstoreDB) CreateObject(input []string) (err error) {

	client := db.client

	ctx := context.Background()

	var (
		keys []*datastore.Key
		vals []Entity
	)

	for _, obj := range input {
		keys = append(keys, datastore.NameKey(keyKind, obj, nil))
		vals = append(vals, Entity{Value: obj})
	}

	if keys, err := client.PutMulti(ctx, keys, vals); err != nil {
		log.Println(err)
	} else {
		for _, key := range keys {
			log.Println("Successfully created a new object : " + key.String())
		}

	}

	return
}

// ReadObject retrieves all the objects for project from Google Cloud Datastore
func (db *CloudstoreDB) ReadObject() (ret map[string]interface{}, err error) {

	var (
		vals   []*Entity
		output []interface{}
	)

	client := db.client

	ctx := context.Background()

	keys, err := client.GetAll(ctx, datastore.NewQuery(keyKind), &vals)
	for _, key := range keys {
		val := strings.Split(key.String(), ",")[1]
		output = append(output, val)
	}

	ret = make(map[string]interface{})
	ret["kind"] = keyKind
	ret["keys"] = output

	return
}
