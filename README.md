# flexstore
A small application with Google Cloud Datastore

The application assumes Go programming enviornment already setup with Go 1.8+
Once your GOROOT and GOPATH is set properly,

## Clone the repository 
git clone https://github.com/sauravhaloi/flexstore.git

## Pre-requisites
Install the following dependencies:
```
go get -u github.com/gorilla/mux
go get -u google.golang.org/appengine
go get -u cloud.google.com/go/datastore
```
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


## Examples

```
$ curl -iX GET "http://localhost:8080/save?input=google&input=cloud&input=datastore"

HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 08 Oct 2017 07:26:04 GMT
Content-Length: 0

$ curl -iX GET "http://localhost:8080/retrieve"

HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 08 Oct 2017 07:55:31 GMT
Content-Length: 111

{"keys":["123","456","666","abc","cloud","datastore","def","google","haloi","saurav","xyz"],"kind":"flexstore"}

$ curl -iX GET "http://localhost:8080/"

HTTP/1.1 200 OK
Date: Sun, 08 Oct 2017 07:55:02 GMT
Content-Length: 12
Content-Type: text/plain; charset=utf-8

Hello world!

$ curl -iX GET "http://localhost:8080/liveness_check"

HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 08 Oct 2017 07:54:54 GMT
Content-Length: 38

{"status":"Service is Up and Running"}
```
