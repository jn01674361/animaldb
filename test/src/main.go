package main

import (
	"fmt"
	"net/http"
)

// This will be the entry point for _GAE_, we'e using GO default http package to pass the request to a handler func
func init() {
	http.HandleFunc("/", handler)
}

// This will handle requests on the default path and return Hello World on the response
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!!!")
}