# flexstore
A small application with Google Cloud Datastore

The application assumes Go programming enviornment already setup with Go 1.8+
Once your GOROOT and GOPATH is set properly,

## Clone the repository 
git clone https://github.com/sauravhaloi/flexstore.git

## Pre-requisites
Install the following dependencies:

go get -u github.com/gorilla/mux

go get -u google.golang.org/appengine

go get -u cloud.google.com/go/datastore

## Building 

$ cd $GOPATH/src/github.com/sauravhaloi/flexstore

$ go build -o flexstore *.go

## Executing
The application runs by default at port 8080

$ ./flexstore

2017/10/08 13:12:40 Successfully connected to Google Cloud Datastore for project xxx-yyy-zzz

2017/10/08 13:12:40 Flexstore is now serving RESTfully [port :8080]...

## APIs

#### GET /save?input=abc 
  - this will save the values of 'input' param in google cloud datastore 
  - multiple values are supported with save?input=input1&input=input2 ... 

#### GET /retrieve - this will retrieve/return all the values saved so far

#### GET /liveness_check - this is a GAE specific application liveness checker endpoint
