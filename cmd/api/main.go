package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/waditya/go-mux-basic-rest-api-project/internal/store"
)

var Products []store.Product // Declare a slice of product

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit : homePage")
	_, err := fmt.Fprintf(w, "<h1>Welcome to the HomePage<h1><br><h3>Did you like it ?</h3>")

	if err != nil {
		fmt.Println("Err writing to the io writer\t", err)
		return
	}
	fmt.Println("Name of the endpoint hit : HomePage") // Write to the console
	return
}

func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit : getAllProducts")
	json.NewEncoder(w).Encode(Products)
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint hit : getProduct , URL Path : ", r.URL.Path)
	key, err := strconv.Atoi(r.URL.Path[len("/product/"):])

	if err != nil {
		log.Println(err)
		return
	}

	for _, product := range Products {
		if int(product.Id) == key {
			json.NewEncoder(w).Encode(product)
		}
	}

}

func requestsHandler() {
	// The Server must serve on some endpoint. Endpoint states the resource that has to be accessed
	// Endpoint could be homepage of the server
	// Use HandleFunc will register the handler function for the given pattern
	// handler function handles the requests coming at a given endpoint
	// In nut-shell, Handler Func ties request to a specific endpoint to a particular Handler Function
	http.HandleFunc("/products", getAllProductsHandler)
	http.HandleFunc("/product/", getProductHandler)
	http.HandleFunc("/", homePageHandler) // homepageHandler is the handler function

	http.ListenAndServe("localhost:10000", nil)

}

func main() {

	// Initialize a slice of Products

	Products = []store.Product{
		store.Product{Id: 1, Name: "Chair", Quantity: 100, Price: 100},
		store.Product{Id: 2, Name: "Table", Quantity: 25, Price: 400},
	}

	requestsHandler()
}
