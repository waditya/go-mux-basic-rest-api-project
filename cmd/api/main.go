package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/waditya/go-mux-basic-rest-api-project/internal/database"
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
	vars := mux.Vars(r)
	key, err := strconv.Atoi(vars["id"])

	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Endpoint hit : getProduct , URL Path : ", r.URL.Path)

	// Below code is for default Mux -- ServeMux
	/* key, err := strconv.Atoi(r.URL.Path[len("/product/"):])

	if err != nil {
		log.Println(err)
		return
	} */

	for _, product := range Products {
		if int(product.Id) == key {
			json.NewEncoder(w).Encode(product)
		} // Add a else case for product not found
	}

}

func requestsHandler() {
	// The Server must serve on some endpoint. Endpoint states the resource that has to be accessed
	// Endpoint could be homepage of the server
	// Use HandleFunc will register the handler function for the given pattern
	// handler function handles the requests coming at a given endpoint
	// In nut-shell, Handler Func ties request to a specific endpoint to a particular Handler Function

	// Create a new instance of mux Router
	muxRouter := mux.NewRouter().StrictSlash(true) // StrictSlash defines the trailing slash behavior for new routes. The initial value is false.

	muxRouter.HandleFunc("/products", getAllProductsHandler)
	muxRouter.HandleFunc("/product/{id}", getProductHandler)
	muxRouter.HandleFunc("/", homePageHandler) // homepageHandler is the handler function

	// http.HandleFunc("/products", getAllProductsHandler)
	// http.HandleFunc("/product/", getProductHandler)
	// http.HandleFunc("/", homePageHandler) // homepageHandler is the handler function

	http.ListenAndServe("localhost:10000", muxRouter)

}

func main() {

	// Initialize a slice of Products

	Products = []store.Product{
		store.Product{Id: 1, Name: "Chair", Quantity: 100, Price: 100},
		store.Product{Id: 2, Name: "Table", Quantity: 25, Price: 400},
	}

	database.NewMySql()

	requestsHandler()

}
