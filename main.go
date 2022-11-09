package main

import (
	"crud-rest-api/controllers"
	"crud-rest-api/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	LoadAppConfig()

	// Initialize Database
	database.Connect(AppConfig.ConnectionString)
	database.Migrate() // call the package and get the exported function

	router := mux.NewRouter().StrictSlash(true) // Initialize the router
	RegisterProductRoutes(router)

	log.Println(fmt.Sprintf("Starting server on port %s", AppConfig.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), router))
}

func RegisterProductRoutes(router *mux.Router) {
	router.HandleFunc("/api/products", controllers.GetProducts).Methods("GET")

	router.HandleFunc("/api/products/{id}", controllers.GetProductById).Methods("GET")

	router.HandleFunc("/api/products", controllers.CreateProduct).Methods("POST")
	router.HandleFunc("/api/products/{id}", controllers.UpdateProduct).Methods("PUT")
	router.HandleFunc("/api/products/{id}", controllers.DeleteProduct).Methods("DELETE")
}
