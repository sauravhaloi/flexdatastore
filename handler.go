package main

import (
	"context"
	"net/http"
	"net/url"
)

// DB is an interface for CRUD operations of flexstore
type DB interface {
	CreateObject(obj url.Values) (err error)
	ReadObject() (obj []map[string]interface{}, err error)
}

// PostValHandler save the values of 'input' param in google cloud datastore
func PostValHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, db *CloudstoreDB) {

}

// GetValHandler retrieve/return all the values saved so far
func GetValHandler(ctx context.Context, w http.ResponseWriter, r *http.Request, db *CloudstoreDB) {

}
