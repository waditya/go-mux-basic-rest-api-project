package main

import (
	"net/http"

	"github.com/waditya/go-mux-basic-rest-api-project/store"
)

func main() {

	// Initialize a slice of Products

	products := []store.Product{
		Product{Id: 1, Name: "Chair", Quantity: 100, Price: 100},
		Product{Id: 2, Name: "Table", Quantity: 25, Price: 400},
	}
	// The Server must serve on some endpoint. Endpoint states the resource that has to be accessed
	// Endpoint could be homepage of the server
	// Use HandleFunc will register the handler function for the given pattern
	// handler function handles the requests coming at a given endpoint
	// In nut-shell, Handler Func ties request to a specific endpoint to a particular Handler Function

	http.HandleFunc("/", homepage) // homepage is the handler function

	http.ListenAndServe("localhost:10000", nil)
}
